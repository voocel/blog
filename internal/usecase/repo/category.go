package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (c CategoryRepo) AddCategoryRepo(ctx context.Context, category *entity.Category) error {
	return c.db.WithContext(ctx).Create(category).Error
}

func (c CategoryRepo) GetCategoryByIdRepo(ctx context.Context, cid int64) (*entity.Category, error) {
	var category = new(entity.Category)
	err := c.db.WithContext(ctx).Where("id = ?", cid).First(category).Error
	return category, err
}

func (c CategoryRepo) GetCategoryByNameRepo(ctx context.Context, name string) (*entity.Category, error) {
	var category = new(entity.Category)
	err := c.db.WithContext(ctx).Where("name = ?", name).First(category).Error
	return category, err
}

func (c CategoryRepo) GetCategoryByNameExistRepo(ctx context.Context, name string) (bool, error) {
	c.db.WithContext(ctx).Where("name = ?", name).First(&entity.Category{})
	return c.db.RowsAffected > 0, c.db.Error
}

func (c CategoryRepo) GetCategoriesRepo(ctx context.Context) ([]*entity.Category, error) {
	categories := make([]*entity.Category, 0)
	err := c.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}
