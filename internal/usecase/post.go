package usecase

import (
	"blog/internal/entity"
	"context"
	"fmt"
	"math"
	"strings"
	"time"
)

type PostUseCase struct {
	postRepo     PostRepo
	categoryRepo CategoryRepo
	tagRepo      TagRepo
}

func NewPostUseCase(postRepo PostRepo, categoryRepo CategoryRepo, tagRepo TagRepo) *PostUseCase {
	return &PostUseCase{
		postRepo:     postRepo,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

func (uc *PostUseCase) Create(ctx context.Context, req entity.CreatePostRequest, author string) error {
	// 设置默认状态
	if req.Status == "" {
		req.Status = "draft"
	}

	post := &entity.Post{
		Title:      req.Title,
		Excerpt:    req.Excerpt,
		Content:    req.Content,
		Author:     author,
		Date:       time.Now().Format("2006-01-02"),
		CategoryID: req.CategoryID,
		ImageUrl:   req.ImageUrl,
		Status:     req.Status,
		Views:      0,
	}

	// 创建文章
	if err := uc.postRepo.Create(ctx, post); err != nil {
		return err
	}

	// 添加标签关联
	if len(req.Tags) > 0 {
		if err := uc.postRepo.AddTags(ctx, post.ID, req.Tags); err != nil {
			return err
		}
	}

	// 更新分类计数
	uc.categoryRepo.IncrementCount(ctx, req.CategoryID)

	return nil
}

func (uc *PostUseCase) GetByID(ctx context.Context, id string) (*entity.PostResponse, error) {
	post, err := uc.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 增加浏览量
	go uc.postRepo.IncrementViews(context.Background(), id)

	return uc.assemblePostResponse(ctx, post)
}

func (uc *PostUseCase) List(ctx context.Context, filters map[string]interface{}, page, limit int) (interface{}, error) {
	posts, total, err := uc.postRepo.List(ctx, filters, page, limit)
	if err != nil {
		return nil, err
	}

	responses := make([]entity.PostResponse, 0, len(posts))
	for _, post := range posts {
		resp, err := uc.assemblePostResponse(ctx, &post)
		if err != nil {
			continue
		}
		responses = append(responses, *resp)
	}

	// 如果有分页参数，返回分页响应
	if page > 0 && limit > 0 {
		totalPages := int(math.Ceil(float64(total) / float64(limit)))
		return entity.PaginatedPostsResponse{
			Data: responses,
			Pagination: entity.Pagination{
				Total:      int(total),
				Page:       page,
				Limit:      limit,
				TotalPages: totalPages,
			},
		}, nil
	}

	// 否则直接返回数组
	return responses, nil
}

func (uc *PostUseCase) Update(ctx context.Context, id string, req entity.UpdatePostRequest) error {
	post, err := uc.postRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// 更新字段
	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Excerpt != "" {
		post.Excerpt = req.Excerpt
	}
	if req.Content != "" {
		post.Content = req.Content
	}
	if req.ImageUrl != "" {
		post.ImageUrl = req.ImageUrl
	}
	if req.Status != "" {
		post.Status = req.Status
	}

	// 处理分类变更
	if req.CategoryID != "" && req.CategoryID != post.CategoryID {
		uc.categoryRepo.DecrementCount(ctx, post.CategoryID)
		uc.categoryRepo.IncrementCount(ctx, req.CategoryID)
		post.CategoryID = req.CategoryID
	}

	// 更新文章
	if err := uc.postRepo.Update(ctx, post); err != nil {
		return err
	}

	// 更新标签
	if req.Tags != nil {
		uc.postRepo.RemoveTags(ctx, id)
		if len(req.Tags) > 0 {
			uc.postRepo.AddTags(ctx, id, req.Tags)
		}
	}

	return nil
}

func (uc *PostUseCase) Delete(ctx context.Context, id string) error {
	post, err := uc.postRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// 删除文章
	if err := uc.postRepo.Delete(ctx, id); err != nil {
		return err
	}

	// 更新分类计数
	uc.categoryRepo.DecrementCount(ctx, post.CategoryID)

	return nil
}

// 组装文章响应
func (uc *PostUseCase) assemblePostResponse(ctx context.Context, post *entity.Post) (*entity.PostResponse, error) {
	resp := &entity.PostResponse{
		ID:         post.ID,
		Title:      post.Title,
		Excerpt:    post.Excerpt,
		Content:    post.Content,
		Author:     post.Author,
		Date:       post.Date,
		CategoryID: post.CategoryID,
		ImageUrl:   post.ImageUrl,
		Views:      post.Views,
		Status:     post.Status,
		ReadTime:   calculateReadTime(post.Content),
		Tags:       []string{},
	}

	// 获取分类名称
	if category, err := uc.categoryRepo.GetByID(ctx, post.CategoryID); err == nil {
		resp.Category = category.Name
	}

	// 获取标签名称
	if tagIDs, err := uc.postRepo.GetTagIDs(ctx, post.ID); err == nil && len(tagIDs) > 0 {
		if tags, err := uc.tagRepo.GetByIDs(ctx, tagIDs); err == nil {
			for _, tag := range tags {
				resp.Tags = append(resp.Tags, tag.Name)
			}
		}
	}

	return resp, nil
}

// 计算阅读时间（约200字/分钟）
func calculateReadTime(content string) string {
	wordCount := len(strings.Fields(content))
	minutes := wordCount / 200
	if minutes < 1 {
		return "1 min read"
	}
	return fmt.Sprintf("%d min read", minutes)
}
