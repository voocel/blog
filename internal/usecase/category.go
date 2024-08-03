package usecase

import (
	"blog/internal/entity"
	"github.com/gin-gonic/gin"
)

type CategoryUseCase struct {
	repo CategoryRepo
}

func NewCategoryUseCase(repo CategoryRepo) *CategoryUseCase {
	return &CategoryUseCase{repo: repo}
}

func (c *CategoryUseCase) List(ctx *gin.Context) ([]*entity.Category, error) {
	return c.repo.GetCategoriesRepo(ctx)
}
