package usecase

import (
	"blog/internal/entity"
	"context"
	"github.com/jinzhu/copier"
)

type BannerUseCase struct {
	repo BannerRepo
}

func NewBannerUseCase(repo BannerRepo) *BannerUseCase {
	return &BannerUseCase{repo: repo}
}

func (c *BannerUseCase) AddBanner(ctx context.Context, req *entity.Banner) error {
	return c.repo.AddBannerRepo(ctx, req)
}

func (c *BannerUseCase) DeleteBanner(ctx context.Context, id int64) error {
	return c.repo.DeleteBannerRepo(ctx, id)
}

func (c *BannerUseCase) DeleteBannerBatch(ctx context.Context, ids []int64) error {
	return c.repo.DeleteBannersBatchRepo(ctx, ids)
}

func (c *BannerUseCase) Detail(ctx context.Context, id int64) (*entity.Banner, error) {
	return c.repo.GetBannerByIdRepo(ctx, id)
}

func (c *BannerUseCase) List(ctx context.Context) ([]*entity.Banner, error) {
	return c.repo.GetBannersRepo(ctx)
}

func (c *BannerUseCase) UpdateBanner(ctx context.Context, req *entity.BannerUpdateReq) error {
	banner := new(entity.Banner)
	copier.Copy(banner, req)
	return c.repo.UpdateBannerRepo(ctx, banner)
}
