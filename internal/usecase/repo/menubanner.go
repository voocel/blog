package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

type MenuBannerRepo struct {
	db *gorm.DB
}

func NewMenuBannerRepo(db *gorm.DB) *MenuBannerRepo {
	return &MenuBannerRepo{db: db}
}

func (m MenuBannerRepo) AddMenuBannerRepo(ctx context.Context, menuBanner *entity.MenuBanner) error {
	//TODO implement me
	panic("implement me")
}

func (m MenuBannerRepo) AddMenuBannerBatchRepo(ctx context.Context, menuBanners []*entity.MenuBanner) error {
	return m.db.Create(menuBanners).Error
}

func (m MenuBannerRepo) GetMenuBannerByIdRepo(ctx context.Context, id int64) (*entity.MenuBanner, error) {
	//TODO implement me
	panic("implement me")
}

func (m MenuBannerRepo) GetMenuBannerByMenuIdRepo(ctx context.Context, menuId int64) (*entity.MenuBanner, error) {
	//TODO implement me
	panic("implement me")
}

func (m MenuBannerRepo) GetMenuBannerByBannerIdRepo(ctx context.Context, bannerId int64) (*entity.MenuBanner, error) {
	//TODO implement me
	panic("implement me")
}

func (m MenuBannerRepo) GetMenuBannersRepo(ctx context.Context) ([]*entity.MenuBanner, error) {
	//TODO implement me
	panic("implement me")
}

func (m MenuBannerRepo) UpdateMenuBannerRepo(ctx context.Context, menuBanner *entity.MenuBanner) error {
	//TODO implement me
	panic("implement me")
}

func (m MenuBannerRepo) DeleteMenuBannerRepo(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
