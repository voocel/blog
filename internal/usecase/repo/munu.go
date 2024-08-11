package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

type MenuRepo struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) *MenuRepo {
	return &MenuRepo{db: db}
}

func (m MenuRepo) AddMenuRepo(ctx context.Context, menu *entity.Menu) (*entity.Menu, error) {
	return menu, m.db.WithContext(ctx).Create(menu).Error
}

func (m MenuRepo) GetMenuByIdRepo(ctx context.Context, id int64) (*entity.Menu, error) {
	menu := new(entity.Menu)
	err := m.db.WithContext(ctx).Where("id = ?", id).First(menu).Error
	return menu, err
}

func (m MenuRepo) GetMenuByPathRepo(ctx context.Context, path string) (*entity.Menu, error) {
	menu := new(entity.Menu)
	err := m.db.WithContext(ctx).Where("path = ?", path).First(menu).Error
	return menu, err
}

func (m MenuRepo) GetMenusRepo(ctx context.Context) ([]*entity.Menu, error) {
	menus := make([]*entity.Menu, 0)
	err := m.db.WithContext(ctx).Find(&menus).Error
	return menus, err
}

func (m MenuRepo) UpdateMenuRepo(ctx context.Context, menu *entity.Menu) error {
	return m.db.WithContext(ctx).Updates(menu).Error
}

func (m MenuRepo) DeleteMenuRepo(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Menu{}).Error
}

func (m MenuRepo) DeleteMenusBatchRepo(ctx context.Context, ids []int64) error {
	return m.db.WithContext(ctx).Where("id in (?)", ids).Delete(&entity.Menu{}).Error
}

func (m MenuRepo) IsTitlePathExistRepo(ctx context.Context, title, path string) bool {
	return m.db.WithContext(ctx).Where("title = ? or path = ?", title, path).First(&entity.Menu{}).RowsAffected > 0
}
