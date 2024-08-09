package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

type StarRepo struct {
	db *gorm.DB
}

func NewStarRepo(db *gorm.DB) *StarRepo {
	return &StarRepo{db: db}
}

func (s StarRepo) AddStarRepo(ctx context.Context, uid, articleId int64) error {
	return s.db.Create(&entity.Star{}).Error
}

func (s StarRepo) DeleteStarRepo(ctx context.Context, uid, articleId int64) error {
	return s.db.WithContext(ctx).Where("user_id = ? and article_id = ?", uid, articleId).Delete(&entity.Star{}).Error
}

func (s StarRepo) GetStarsByUidRepo(ctx context.Context, uid int64) ([]*entity.Star, error) {
	var stars []*entity.Star
	return stars, s.db.WithContext(ctx).Where("user_id = ?", uid).Find(&stars).Error
}

func (s StarRepo) GetStarsByArticleIdRepo(ctx context.Context, articleId int64) ([]*entity.Star, error) {
	var stars []*entity.Star
	return stars, s.db.WithContext(ctx).Where("article_id = ?", articleId).Find(&stars).Error
}
