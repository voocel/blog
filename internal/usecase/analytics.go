package usecase

import (
	"blog/internal/entity"
	"blog/pkg/geoip"
	"context"
	"time"
)

type AnalyticsUseCase struct {
	analyticsRepo AnalyticsRepo
}

func NewAnalyticsUseCase(analyticsRepo AnalyticsRepo) *AnalyticsUseCase {
	return &AnalyticsUseCase{analyticsRepo: analyticsRepo}
}

func (uc *AnalyticsUseCase) LogVisit(ctx context.Context, req entity.LogVisitRequest, ip, userAgent string) error {
	// 使用 GeoIP 查询地理位置
	location := geoip.Lookup(ip)

	// 如果未提供文章 ID，保存为 NULL 避免空字符串触发 UUID 解析错误
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
