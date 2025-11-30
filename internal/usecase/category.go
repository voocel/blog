package usecase

import (
	"blog/internal/entity"
	"context"
	"strings"
)

type CategoryUseCase struct {
	categoryRepo CategoryRepo
}

func NewCategoryUseCase(categoryRepo CategoryRepo) *CategoryUseCase {
	return &CategoryUseCase{categoryRepo: categoryRepo}
}

func (uc *CategoryUseCase) Create(ctx context.Context, req entity.CreateCategoryRequest) error {
	// Auto-generate slug if empty
	slug := req.Slug
	if slug == "" {
		slug = generateSlug(req.Name)
	}

	category := &entity.Category{
		Name:  req.Name,
		Slug:  slug,
		Count: 0,
	}

	return uc.categoryRepo.Create(ctx, category)
}

func (uc *CategoryUseCase) List(ctx context.Context) ([]entity.CategoryResponse, error) {
	categories, err := uc.categoryRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]entity.CategoryResponse, len(categories))
	for i, cat := range categories {
		responses[i] = entity.CategoryResponse{
			ID:    cat.ID,
			Name:  cat.Name,
			Slug:  cat.Slug,
			Count: cat.Count,
		}
	}

	return responses, nil
}

func (uc *CategoryUseCase) Delete(ctx context.Context, id string) error {
	return uc.categoryRepo.Delete(ctx, id)
}

// generateSlug generates slug: convert to lowercase, replace spaces with -
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
