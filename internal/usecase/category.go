package usecase

import (
	"blog/internal/entity"
	"context"
	"errors"
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

func (c *CategoryUseCase) Create(ctx context.Context, req *entity.CategoryRequest) error {
	// 检查名称是否已存在
	exists, err := c.repo.GetCategoryByNameExistRepo(ctx, req.Name)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("分类名称已存在")
	}

	// 检查路径是否已存在
	exists, err = c.repo.GetCategoryByPathExistRepo(ctx, req.Path)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("分类路径已存在")
	}

	category := &entity.Category{
		Name:        req.Name,
		Path:        req.Path,
		Description: req.Description,
	}

	return c.repo.AddCategoryRepo(ctx, category)
}

func (c *CategoryUseCase) Update(ctx context.Context, id int64, req *entity.CategoryRequest) error {
	// 检查分类是否存在
	category, err := c.repo.GetCategoryByIdRepo(ctx, id)
	if err != nil {
		return errors.New("分类不存在")
	}

	// 检查名称是否被其他分类使用
	if category.Name != req.Name {
		exists, err := c.repo.GetCategoryByNameExistRepo(ctx, req.Name)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("分类名称已存在")
		}
	}

	// 检查路径是否被其他分类使用
	if category.Path != req.Path {
		exists, err := c.repo.GetCategoryByPathExistRepo(ctx, req.Path)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("分类路径已存在")
		}
	}

	// 更新分类信息
	category.Name = req.Name
	category.Path = req.Path
	category.Description = req.Description

	return c.repo.UpdateCategoryRepo(ctx, category)
}

func (c *CategoryUseCase) Delete(ctx context.Context, id int64) error {
	// 检查分类是否存在
	_, err := c.repo.GetCategoryByIdRepo(ctx, id)
	if err != nil {
		return errors.New("分类不存在")
	}

	// TODO: 检查是否有文章使用该分类
	// 可以在这里添加检查逻辑，防止删除正在使用的分类

	return c.repo.DeleteCategoryRepo(ctx, id)
}

func (c *CategoryUseCase) GetByID(ctx context.Context, id int64) (*entity.Category, error) {
	return c.repo.GetCategoryByIdRepo(ctx, id)
}
