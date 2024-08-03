package usecase

import (
	"blog/internal/entity"
	"context"
)

type CategoryUseCase struct {
	repo CategoryRepo
}

func NewCategoryUseCase(repo CategoryRepo) *CategoryUseCase {
	return &CategoryUseCase{repo: repo}
}

func (c *CategoryUseCase) List(ctx context.Context) ([]*entity.Category, error) {
	return c.repo.GetCategoriesRepo(ctx)
}
