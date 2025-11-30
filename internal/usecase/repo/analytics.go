package repo

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"context"
	"time"

	"gorm.io/gorm"
)

type analyticsRepo struct {
	db *gorm.DB
}

func NewAnalyticsRepo(db *gorm.DB) usecase.AnalyticsRepo {
	return &analyticsRepo{db: db}
}

func (r *analyticsRepo) Create(ctx context.Context, log *entity.Analytics) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *analyticsRepo) GetLogs(ctx context.Context, startDate, endDate string, limit int) ([]entity.Analytics, error) {
	var logs []entity.Analytics
	query := r.db.WithContext(ctx).Model(&entity.Analytics{})

	if startDate != "" {
		start, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			query = query.Where("timestamp >= ?", start.Unix())
		}
	}

	if endDate != "" {
		end, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			query = query.Where("timestamp <= ?", end.Unix()+86400)
		}
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Order("timestamp DESC").Find(&logs).Error
	return logs, err
}
