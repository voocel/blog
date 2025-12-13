package usecase

import (
	"blog/internal/entity"
	"blog/pkg/log"
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
	if strings.TrimSpace(req.Excerpt) == "" {
		req.Excerpt = deriveExcerpt(req.Content, 180)
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

	if err := uc.postRepo.CreateWithTags(ctx, post, req.Tags); err != nil {
		return err
	}

	if err := uc.categoryRepo.IncrementCount(ctx, req.CategoryID); err != nil {
		// Non-critical counter update: log and continue.
		log.Warnw("Increment category count failed",
			log.Pair("category_id", req.CategoryID),
			log.Pair("error", err.Error()),
		)
	}

	return nil
}

func deriveExcerpt(content string, maxRunes int) string {
	// Best-effort excerpt: collapse whitespace, then truncate by runes.
	s := strings.TrimSpace(content)
	if s == "" {
		return ""
	}
	s = strings.Join(strings.Fields(s), " ")
	if maxRunes <= 0 {
		return s
	}
	r := []rune(s)
	if len(r) <= maxRunes {
		return s
	}
	return strings.TrimSpace(string(r[:maxRunes])) + "…"
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

	responses, err := uc.assemblePostResponsesBatch(ctx, posts)
	if err != nil {
		return nil, err
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

func (uc *PostUseCase) assemblePostResponsesBatch(ctx context.Context, posts []entity.Post) ([]entity.PostResponse, error) {
	if len(posts) == 0 {
		return []entity.PostResponse{}, nil
	}

	postIDs := make([]string, 0, len(posts))
	categoryIDSet := make(map[string]struct{}, len(posts))
	for i := range posts {
		postIDs = append(postIDs, posts[i].ID)
		if posts[i].CategoryID != "" {
			categoryIDSet[posts[i].CategoryID] = struct{}{}
		}
	}

	// Batch load categories
	categoryIDs := make([]string, 0, len(categoryIDSet))
	for id := range categoryIDSet {
		categoryIDs = append(categoryIDs, id)
	}
	categories, err := uc.categoryRepo.GetByIDs(ctx, categoryIDs)
	if err != nil {
		return nil, err
	}
	categoryNameByID := make(map[string]string, len(categories))
	for i := range categories {
		categoryNameByID[categories[i].ID] = categories[i].Name
	}

	// Batch load post->tagIDs
	tagIDsByPostID, err := uc.postRepo.GetTagIDsByPostIDs(ctx, postIDs)
	if err != nil {
		return nil, err
	}
	tagIDSet := make(map[string]struct{})
	for _, ids := range tagIDsByPostID {
		for _, id := range ids {
			if id != "" {
				tagIDSet[id] = struct{}{}
			}
		}
	}
	allTagIDs := make([]string, 0, len(tagIDSet))
	for id := range tagIDSet {
		allTagIDs = append(allTagIDs, id)
	}

	// Batch load tags
	tagNameByID := make(map[string]string)
	if len(allTagIDs) > 0 {
		tags, err := uc.tagRepo.GetByIDs(ctx, allTagIDs)
		if err != nil {
			return nil, err
		}
		tagNameByID = make(map[string]string, len(tags))
		for i := range tags {
			tagNameByID[tags[i].ID] = tags[i].Name
		}
	}

	responses := make([]entity.PostResponse, 0, len(posts))
	for i := range posts {
		p := posts[i]
		resp := entity.PostResponse{
			ID:         p.ID,
			Title:      p.Title,
			Excerpt:    p.Excerpt,
			Content:    p.Content,
			Author:     p.Author,
			Date:       p.Date,
			CategoryID: p.CategoryID,
			Category:   categoryNameByID[p.CategoryID],
			ReadTime:   calculateReadTime(p.Content),
			ImageUrl:   p.ImageUrl,
			Views:      p.Views,
			Status:     p.Status,
			Tags:       []string{},
		}

		for _, tagID := range tagIDsByPostID[p.ID] {
			if name := tagNameByID[tagID]; name != "" {
				resp.Tags = append(resp.Tags, name)
			}
		}
		responses = append(responses, resp)
	}
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
		if err := uc.categoryRepo.DecrementCount(ctx, post.CategoryID); err != nil {
			log.Warnw("Decrement category count failed",
				log.Pair("category_id", post.CategoryID),
				log.Pair("error", err.Error()),
			)
		}
		if err := uc.categoryRepo.IncrementCount(ctx, req.CategoryID); err != nil {
			log.Warnw("Increment category count failed",
				log.Pair("category_id", req.CategoryID),
				log.Pair("error", err.Error()),
			)
		}
		post.CategoryID = req.CategoryID
	}

	// Update post (and tags optionally)
	if req.Tags != nil {
		if err := uc.postRepo.UpdateWithTags(ctx, post, req.Tags); err != nil {
			return err
		}
	} else {
		if err := uc.postRepo.Update(ctx, post); err != nil {
			return err
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

	if err := uc.categoryRepo.DecrementCount(ctx, post.CategoryID); err != nil {
		log.Warnw("Decrement category count failed",
			log.Pair("category_id", post.CategoryID),
			log.Pair("error", err.Error()),
		)
	}

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
