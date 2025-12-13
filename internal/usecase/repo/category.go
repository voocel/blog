package repo

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"context"
	"errors"

	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) usecase.CategoryRepo {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) Create(ctx context.Context, category *entity.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *categoryRepo) GetByID(ctx context.Context, id string) (*entity.Category, error) {
	var category entity.Category
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepo) GetByIDs(ctx context.Context, ids []string) ([]entity.Category, error) {
	if len(ids) == 0 {
		return []entity.Category{}, nil
	}
	var categories []entity.Category
	if err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepo) GetBySlug(ctx context.Context, slug string) (*entity.Category, error) {
	var category entity.Category
	err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepo) List(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.WithContext(ctx).Order("name ASC").Find(&categories).Error
	return categories, err
}

func (r *categoryRepo) Update(ctx context.Context, category *entity.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}

func (r *categoryRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Category{}).Error
}

func (r *categoryRepo) IncrementCount(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&entity.Category{}).
		Where("id = ?", id).
		UpdateColumn("count", gorm.Expr("count + 1")).Error
}

func (r *categoryRepo) DecrementCount(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&entity.Category{}).
		Where("id = ?", id).
		UpdateColumn("count", gorm.Expr("CASE WHEN count > 0 THEN count - 1 ELSE 0 END")).Error
}

// Count counts total number of categories
func (r *categoryRepo) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Count(&count).Error
	return count, err
}
