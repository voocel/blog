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
	return b.db.WithContext(ctx).Create(banner).Error
}

func (b BannerRepo) GetBannerByIdRepo(ctx context.Context, id int64) (*entity.Banner, error) {
	var banner = new(entity.Banner)
	err := b.db.WithContext(ctx).Where("id = ?", id).First(banner).Error
	return banner, err
}

func (b BannerRepo) GetBannersRepo(ctx context.Context) ([]*entity.Banner, error) {
	banners := make([]*entity.Banner, 0)
	err := b.db.WithContext(ctx).Find(&banners).Error
	return banners, err
}

func (b BannerRepo) UpdateBannerRepo(ctx context.Context, banner *entity.Banner) error {
	return b.db.WithContext(ctx).Updates(banner).Error
}

func (b BannerRepo) DeleteBannerRepo(ctx context.Context, id int64) error {
	return b.db.WithContext(ctx).Delete(&entity.Banner{}, id).Error
}

func (b BannerRepo) DeleteBannersBatchRepo(ctx context.Context, ids []int64) error {
	return b.db.WithContext(ctx).Delete(&entity.Banner{}, ids).Error
}
