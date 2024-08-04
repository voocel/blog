package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

type BannerRepo struct {
	db *gorm.DB
}

func NewBannerRepo(db *gorm.DB) *BannerRepo {
	return &BannerRepo{db: db}
}

func (b BannerRepo) AddBannerRepo(ctx context.Context, banner *entity.Banner) error {
	//TODO implement me
	panic("implement me")
}

func (b BannerRepo) GetBannerByIdRepo(ctx context.Context, id int64) (*entity.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (b BannerRepo) GetBannersRepo(ctx context.Context) ([]*entity.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (b BannerRepo) UpdateBannerRepo(ctx context.Context, banner *entity.Banner) error {
	//TODO implement me
	panic("implement me")
}

func (b BannerRepo) DeleteBannerRepo(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (b BannerRepo) DeleteBannersBatchRepo(ctx context.Context, ids []int64) error {
	//TODO implement me
	panic("implement me")
}
