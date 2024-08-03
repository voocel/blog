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
	//TODO implement me
	panic("implement me")
}

func (s StarRepo) DeleteStarRepo(ctx context.Context, uid, articleId int64) error {
	//TODO implement me
	panic("implement me")
}

func (s StarRepo) GetStarsByUidRepo(ctx context.Context, uid int64) ([]*entity.Star, error) {
	//TODO implement me
	panic("implement me")
}

func (s StarRepo) GetStarsByArticleIdRepo(ctx context.Context, articleId int64) ([]*entity.Star, error) {
	//TODO implement me
	panic("implement me")
}
