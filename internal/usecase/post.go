package usecase

import (
	"blog/internal/entity"
	"blog/pkg/geoip"
	"blog/pkg/log"
	"context"
	"fmt"
	"math"
	"strings"
	"time"
)

type PostUseCase struct {
	postRepo      PostRepo
	categoryRepo  CategoryRepo
	tagRepo       TagRepo
	analyticsRepo AnalyticsRepo
}

func NewPostUseCase(postRepo PostRepo, categoryRepo CategoryRepo, tagRepo TagRepo, analyticsRepo AnalyticsRepo) *PostUseCase {
	return &PostUseCase{
		postRepo:      postRepo,
		categoryRepo:  categoryRepo,
		tagRepo:       tagRepo,
		analyticsRepo: analyticsRepo,
	}
}

func (uc *PostUseCase) Create(ctx context.Context, req entity.CreatePostRequest, author string) error {
	status, err := normalizePostStatus(req.Status)
	if err != nil {
		return err
	}

	if strings.TrimSpace(req.Excerpt) == "" {
		req.Excerpt = deriveExcerpt(req.Content, 180)
	}

	publishAt, err := normalizePublishAt(req.PublishAt, time.Now())
	if err != nil {
		return err
	}

	post := &entity.Post{
		Title:      req.Title,
		Excerpt:    req.Excerpt,
		Content:    req.Content,
		Author:     author,
		PublishAt:  publishAt,
		CategoryID: req.CategoryID,
		Cover:      req.Cover,
		Status:     status,
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

	return uc.assemblePostResponse(ctx, post)
}

// GetByIDWithAnalytics retrieves a post and logs the visit (used by public API)
func (uc *PostUseCase) GetByIDWithAnalytics(ctx context.Context, id, ip, userAgent string) (*entity.PostResponse, error) {
	post, err := uc.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Increment views (async with timeout)
	go func() {
		bgCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		uc.postRepo.IncrementViews(bgCtx, id)
	}()

	// Log visit (async)
	go uc.logVisit(id, post.Title, ip, userAgent)

	return uc.assemblePostResponse(ctx, post)
}

// logVisit records a page visit to analytics
func (uc *PostUseCase) logVisit(postID, postTitle, ip, userAgent string) {
	if uc.analyticsRepo == nil {
		return
	}

	location := geoip.Lookup(ip)

	pID := postID
	log := &entity.Analytics{
		PagePath:  "/post/" + postID,
		PostID:    &pID,
		PostTitle: postTitle,
		IP:        ip,
		Location:  location,
		UserAgent: userAgent,
		Timestamp: time.Now().Unix(),
	}

	bgCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = uc.analyticsRepo.Create(bgCtx, log)
}

// LogHomeVisit records a homepage visit to analytics
func (uc *PostUseCase) LogHomeVisit(ip, userAgent string) {
	if uc.analyticsRepo == nil {
		return
	}

	location := geoip.Lookup(ip)

	log := &entity.Analytics{
		PagePath:  "/",
		IP:        ip,
		Location:  location,
		UserAgent: userAgent,
		Timestamp: time.Now().Unix(),
	}

	bgCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = uc.analyticsRepo.Create(bgCtx, log)
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
			PublishAt:  p.PublishAt,
			CategoryID: p.CategoryID,
			Category:   categoryNameByID[p.CategoryID],
			ReadTime:   calculateReadTime(p.Content),
			Cover:      p.Cover,
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
	if req.Cover != "" {
		post.Cover = req.Cover
	}
	if req.Status != "" {
		status, err := normalizePostStatus(req.Status)
		if err != nil {
			return err
		}
		post.Status = status
	}
	if req.PublishAt != "" {
		publishAt, err := normalizePublishAt(req.PublishAt, time.Now())
		if err != nil {
			return err
		}
		post.PublishAt = publishAt
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
		PublishAt:  post.PublishAt,
		CategoryID: post.CategoryID,
		Cover:      post.Cover,
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

// calculateReadTime calculates reading time based on content type.
// Chinese: ~500 characters/minute, English: ~200 words/minute.
func calculateReadTime(content string) string {
	// Count Chinese characters and English words separately
	var chineseChars, englishWords int
	inWord := false

	for _, r := range content {
		if r >= 0x4E00 && r <= 0x9FFF { // CJK Unified Ideographs
			chineseChars++
			inWord = false
		} else if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			if !inWord {
				englishWords++
				inWord = true
			}
		} else {
			inWord = false
		}
	}

	// Calculate reading time: Chinese 500 chars/min, English 200 words/min
	minutes := float64(chineseChars)/500 + float64(englishWords)/200

	if minutes < 1 {
		return "1 min read"
	}
	return fmt.Sprintf("%d min read", int(minutes+0.5)) // Round to nearest
}

func normalizePostStatus(s string) (string, error) {
	v := strings.ToLower(strings.TrimSpace(s))
	if v == "" {
		return "draft", nil
	}
	switch v {
	case "draft", "published":
		return v, nil
	default:
		return "", fmt.Errorf("%w: invalid status: must be 'draft' or 'published'", ErrInvalidArgument)
	}
}

// normalizePublishAt returns a concrete scheduled publish time.
// It accepts RFC3339 / RFC3339Nano strings; if omitted, defaults to now.
func normalizePublishAt(input string, now time.Time) (time.Time, error) {
	s := strings.TrimSpace(input)
	if s == "" {
		return now, nil
	}

	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return t, nil
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t, nil
	}
	return time.Time{}, fmt.Errorf("%w: invalid publishAt: expected RFC3339 time, e.g. 2025-12-14T16:30:00+08:00", ErrInvalidArgument)
}
