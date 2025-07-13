package repo

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"context"

	"gorm.io/gorm"
)

type statisticsRepo struct {
	db *gorm.DB
}

func NewStatisticsRepo(db *gorm.DB) usecase.StatisticsRepo {
	return &statisticsRepo{
		db: db,
	}
}

func (r *statisticsRepo) GetUsersCountRepo(ctx context.Context) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.User{}).Count(&count).Error
	return int(count), err
}

func (r *statisticsRepo) GetArticlesCountRepo(ctx context.Context) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Article{}).Count(&count).Error
	return int(count), err
}

func (r *statisticsRepo) GetCommentsCountRepo(ctx context.Context) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Comment{}).Count(&count).Error
	return int(count), err
}

func (r *statisticsRepo) GetVisitsCountRepo(ctx context.Context) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Article{}).Select("COALESCE(SUM(view_count), 0)").Scan(&count).Error
	return int(count), err
}
