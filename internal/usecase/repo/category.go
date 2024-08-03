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
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepo) GetCategoryByIdRepo(ctx context.Context, cid int64) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepo) GetCategoryByNameRepo(ctx context.Context, name string) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepo) GetCategoryByNameExistRepo(ctx context.Context, name string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepo) GetCategoriesRepo(ctx context.Context) ([]*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}
