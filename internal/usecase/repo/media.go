package repo

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"context"
	"errors"

	"gorm.io/gorm"
)

type mediaRepo struct {
	db *gorm.DB
}

func NewMediaRepo(db *gorm.DB) usecase.MediaRepo {
	return &mediaRepo{db: db}
}

func (r *mediaRepo) Create(ctx context.Context, media *entity.Media) error {
	return r.db.WithContext(ctx).Create(media).Error
}

func (r *mediaRepo) GetByID(ctx context.Context, id string) (*entity.Media, error) {
	var media entity.Media
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&media).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("media not found")
		}
		return nil, err
	}
	return &media, nil
}

func (r *mediaRepo) List(ctx context.Context) ([]entity.Media, error) {
	var media []entity.Media
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&media).Error
	return media, err
}

func (r *mediaRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Media{}).Error
}

// Count counts total number of media files
func (r *mediaRepo) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Media{}).Count(&count).Error
	return count, err
}
