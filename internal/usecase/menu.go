package usecase

import (
	"blog/internal/entity"
	"context"
	"github.com/jinzhu/copier"
)

type MenuUseCase struct {
	repo MenuRepo
}

func NewMenuUseCase(repo MenuRepo) *MenuUseCase {
	return &MenuUseCase{repo: repo}
}

func (c *MenuUseCase) AddMenu(ctx context.Context, req entity.MenuReq) (*entity.Menu, error) {
	menu := new(entity.Menu)
	copier.Copy(menu, req)
	return c.repo.AddMenuRepo(ctx, menu)
}

func (c *MenuUseCase) Detail(ctx context.Context, id int64) (*entity.Menu, error) {
	return c.repo.GetMenuByIdRepo(ctx, id)
}

func (c *MenuUseCase) DetailByPath(ctx context.Context, path string) (*entity.Menu, error) {
	return c.repo.GetMenuByPathRepo(ctx, path)
}

func (c *MenuUseCase) List(ctx context.Context) ([]*entity.Menu, error) {
	return c.repo.GetMenusRepo(ctx)
}

func (c *MenuUseCase) UpdateMenu(ctx context.Context, req entity.MenuReq) error {
	menu := new(entity.Menu)
	copier.Copy(menu, req)
	return c.repo.UpdateMenuRepo(ctx, menu)
}

func (c *MenuUseCase) DeleteMenu(ctx context.Context, id int64) error {
	return c.repo.DeleteMenuRepo(ctx, id)
}

func (c *MenuUseCase) DeleteMenusBatch(ctx context.Context, ids []int64) error {
	return c.repo.DeleteMenusBatchRepo(ctx, ids)
}

func (c *MenuUseCase) IsTitlePathExist(ctx context.Context, title, path string) bool {
	return c.repo.IsTitlePathExistRepo(ctx, title, path)
}
