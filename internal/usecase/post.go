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
	if req.Status == "" {
		req.Status = "draft"
	}

	// Use provided date or default to current time
	date := req.Date
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	post := &entity.Post{
		Title:      req.Title,
		Excerpt:    req.Excerpt,
		Content:    req.Content,
		Author:     author,
		Date:       date,
		CategoryID: req.CategoryID,
		ImageUrl:   req.ImageUrl,
		Status:     req.Status,
		Views:      0,
	}

	if err := uc.postRepo.Create(ctx, post); err != nil {
		return err
	}

	if len(req.Tags) > 0 {
		if err := uc.postRepo.AddTags(ctx, post.ID, req.Tags); err != nil {
			return err
		}
	}

	uc.categoryRepo.IncrementCount(ctx, req.CategoryID)

	return nil
}

func (uc *PostUseCase) GetByID(ctx context.Context, id string) (*entity.PostResponse, error) {
	post, err := uc.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Increment views
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

	// Return paginated response if pagination parameters provided
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

	// Otherwise return array directly
	return responses, nil
}

func (uc *PostUseCase) Update(ctx context.Context, id string, req entity.UpdatePostRequest) error {
	post, err := uc.postRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Update fields
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
	if req.Date != "" {
		post.Date = req.Date
	}

	// Handle category change
	if req.CategoryID != "" && req.CategoryID != post.CategoryID {
		uc.categoryRepo.DecrementCount(ctx, post.CategoryID)
		uc.categoryRepo.IncrementCount(ctx, req.CategoryID)
		post.CategoryID = req.CategoryID
	}

	if err := uc.postRepo.Update(ctx, post); err != nil {
		return err
	}

	// Update tags
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

	if err := uc.postRepo.Delete(ctx, id); err != nil {
		return err
	}

	uc.categoryRepo.DecrementCount(ctx, post.CategoryID)

	return nil
}

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

	if category, err := uc.categoryRepo.GetByID(ctx, post.CategoryID); err == nil {
		resp.Category = category.Name
	}

	if tagIDs, err := uc.postRepo.GetTagIDs(ctx, post.ID); err == nil && len(tagIDs) > 0 {
		if tags, err := uc.tagRepo.GetByIDs(ctx, tagIDs); err == nil {
			for _, tag := range tags {
				resp.Tags = append(resp.Tags, tag.Name)
			}
		}
	}

	return resp, nil
}

// calculateReadTime calculates reading time (approximately 200 words/minute)
func calculateReadTime(content string) string {
	wordCount := len(strings.Fields(content))
	minutes := wordCount / 200
	if minutes < 1 {
		return "1 min read"
	}
	return fmt.Sprintf("%d min read", minutes)
}
