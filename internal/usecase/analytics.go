package usecase

import (
	"blog/internal/entity"
	"blog/pkg/geoip"
	"blog/pkg/log"
	"context"
	"time"
)

type AnalyticsUseCase struct {
	analyticsRepo AnalyticsRepo
	postRepo      PostRepo
	categoryRepo  CategoryRepo
	tagRepo       TagRepo
	mediaRepo     MediaRepo
}

func NewAnalyticsUseCase(analyticsRepo AnalyticsRepo, postRepo PostRepo, categoryRepo CategoryRepo, tagRepo TagRepo, mediaRepo MediaRepo) *AnalyticsUseCase {
	return &AnalyticsUseCase{
		analyticsRepo: analyticsRepo,
		postRepo:      postRepo,
		categoryRepo:  categoryRepo,
		tagRepo:       tagRepo,
		mediaRepo:     mediaRepo,
	}
}

func (uc *AnalyticsUseCase) LogVisit(ctx context.Context, req entity.LogVisitRequest, ip, userAgent string) error {
	location := geoip.Lookup(ip)

	var postID *string
	if req.PostID != "" {
		postID = &req.PostID
	}

	log := &entity.Analytics{
		PagePath:  req.PagePath,
		PostID:    postID,
		PostTitle: req.PostTitle,
		IP:        ip,
		Location:  location,
		Timestamp: time.Now().Unix(),
		UserAgent: userAgent,
	}

	return uc.analyticsRepo.Create(ctx, log)
}

func (uc *AnalyticsUseCase) GetLogs(ctx context.Context, startDate, endDate string, limit int) ([]entity.AnalyticsResponse, error) {
	if limit == 0 {
		limit = 100
	}

	logs, err := uc.analyticsRepo.GetLogs(ctx, startDate, endDate, limit)
	if err != nil {
		return nil, err
	}

	responses := make([]entity.AnalyticsResponse, len(logs))
	for i, log := range logs {
		responses[i] = entity.AnalyticsResponse{
			ID:        log.ID,
			PagePath:  log.PagePath,
			PostID:    log.PostID,
			PostTitle: log.PostTitle,
			IP:        log.IP,
			Location:  log.Location,
			Timestamp: log.Timestamp,
			UserAgent: log.UserAgent,
		}
	}

	return responses, nil
}

// GetDashboardOverview retrieves dashboard overview data
func (uc *AnalyticsUseCase) GetDashboardOverview(ctx context.Context) (*entity.DashboardOverviewResponse, error) {
	postsCount, err := uc.postRepo.Count(ctx)
	if err != nil {
		log.Warnf("failed to get posts count: %v", err)
	}
	categoriesCount, err := uc.categoryRepo.Count(ctx)
	if err != nil {
		log.Warnf("failed to get categories count: %v", err)
	}
	tagsCount, err := uc.tagRepo.Count(ctx)
	if err != nil {
		log.Warnf("failed to get tags count: %v", err)
	}
	filesCount, err := uc.mediaRepo.Count(ctx)
	if err != nil {
		log.Warnf("failed to get files count: %v", err)
	}

	recentPosts, err := uc.postRepo.GetRecent(ctx, 3)
	if err != nil {
		log.Warnf("failed to get recent posts: %v", err)
	}

	recentPostResponses := make([]entity.PostResponse, 0, len(recentPosts))
	for _, post := range recentPosts {
		resp := entity.PostResponse{
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

		recentPostResponses = append(recentPostResponses, resp)
	}

	systemStatus := entity.DashboardSystemStatus{
		StorageUsage: 45, // todo
		AIQuota:      60, // todo
	}

	return &entity.DashboardOverviewResponse{
		Counts: entity.DashboardCounts{
			Posts:      postsCount,
			Categories: categoriesCount,
			Tags:       tagsCount,
			Files:      filesCount,
		},
		RecentPosts:  recentPostResponses,
		SystemStatus: systemStatus,
	}, nil
}
