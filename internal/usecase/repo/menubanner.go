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
	return m.db.Create(menuBanner).Error
}

func (m MenuBannerRepo) AddMenuBannerBatchRepo(ctx context.Context, menuBanners []*entity.MenuBanner) error {
	return m.db.Create(menuBanners).Error
}

func (m MenuBannerRepo) GetMenuBannerByIdRepo(ctx context.Context, id int64) (*entity.MenuBanner, error) {
	var menuBanner entity.MenuBanner
	return &menuBanner, m.db.Where("id = ?", id).First(&menuBanner).Error
}

func (m MenuBannerRepo) GetMenuBannerByMenuIdRepo(ctx context.Context, menuId int64) ([]*entity.MenuBanner, error) {
	var menuBanners []*entity.MenuBanner
	return menuBanners, m.db.Where("menu_id = ?", menuId).Find(&menuBanners).Error
}

func (m MenuBannerRepo) GetMenuBannerByBannerIdRepo(ctx context.Context, bannerId int64) ([]*entity.MenuBanner, error) {
	var menuBanners []*entity.MenuBanner
	return menuBanners, m.db.Where("banner_id = ?", bannerId).Find(&menuBanners).Error
}

func (m MenuBannerRepo) GetMenuBannersRepo(ctx context.Context) ([]*entity.MenuBanner, error) {
	var menuBanners []*entity.MenuBanner
	return menuBanners, m.db.Find(&menuBanners).Error
}

func (m MenuBannerRepo) UpdateMenuBannerRepo(ctx context.Context, menuBanner *entity.MenuBanner) error {
	return m.db.Save(menuBanner).Error
}

func (m MenuBannerRepo) DeleteMenuBannerRepo(ctx context.Context, id int64) error {
	return m.db.Where("id = ?", id).Delete(&entity.MenuBanner{}).Error
}
