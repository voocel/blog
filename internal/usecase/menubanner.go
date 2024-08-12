package usecase

import (
	"blog/internal/entity"
	"context"
)

type MenuBannerUseCase struct {
	repo MenuBannerRepo
}

func NewMenuBannerUseCase(repo MenuBannerRepo) *MenuBannerUseCase {
	return &MenuBannerUseCase{repo: repo}
}

func (c *MenuBannerUseCase) AddMenuBanner(ctx context.Context, menuBanner *entity.MenuBanner) error {
	return c.repo.AddMenuBannerRepo(ctx, menuBanner)
}

func (c *MenuBannerUseCase) AddMenuBannerBatch(ctx context.Context, menuBanners []*entity.MenuBanner) error {
	return c.repo.AddMenuBannerBatchRepo(ctx, menuBanners)
}

func (c *MenuBannerUseCase) DeleteMenuBanner(ctx context.Context, id int64) error {
	return c.repo.DeleteMenuBannerRepo(ctx, id)
}

func (c *MenuBannerUseCase) GetMenuBannerByMenuId(ctx context.Context, menuId int64) ([]*entity.MenuBanner, error) {
	return c.repo.GetMenuBannerByMenuIdRepo(ctx, menuId)
}

func (c *MenuBannerUseCase) GetMenuBannerByBannerId(ctx context.Context, bannerId int64) ([]*entity.MenuBanner, error) {
	return c.repo.GetMenuBannerByBannerIdRepo(ctx, bannerId)
}

func (c *MenuBannerUseCase) GetMenuBannerById(ctx context.Context, id int64) (*entity.MenuBanner, error) {
	return c.repo.GetMenuBannerByIdRepo(ctx, id)
}

func (c *MenuBannerUseCase) GetMenuBanners(ctx context.Context) ([]*entity.MenuBanner, error) {
	return c.repo.GetMenuBannersRepo(ctx)
}

func (c *MenuBannerUseCase) UpdateMenuBanner(ctx context.Context, menuBanner *entity.MenuBanner) error {
	return c.repo.UpdateMenuBannerRepo(ctx, menuBanner)
}
