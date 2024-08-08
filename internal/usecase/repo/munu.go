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

func (m MenuRepo) AddMenuRepo(ctx context.Context, menu *entity.Menu) error {
	//TODO implement me
	panic("implement me")
}

func (m MenuRepo) GetMenuByIdRepo(ctx context.Context, id int64) (*entity.Menu, error) {
	//TODO implement me
	panic("implement me")
}

func (m MenuRepo) GetMenuByPathRepo(ctx context.Context, path string) (*entity.Menu, error) {
	//TODO implement me
	panic("implement me")
}

func (m MenuRepo) GetMenusRepo(ctx context.Context) ([]*entity.Menu, error) {
	//TODO implement me
	panic("implement me")
}

func (m MenuRepo) UpdateMenuRepo(ctx context.Context, menu *entity.Menu) error {
	//TODO implement me
	panic("implement me")
}

func (m MenuRepo) DeleteMenuRepo(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MenuRepo) DeleteMenusBatchRepo(ctx context.Context, ids []int64) error {
	//TODO implement me
	panic("implement me")
}
