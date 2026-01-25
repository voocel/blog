package repo

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"context"

	"gorm.io/gorm"
)

type likeRepo struct {
	db *gorm.DB
}

func NewLikeRepo(db *gorm.DB) usecase.LikeRepo {
	return &likeRepo{db: db}
}

// Create adds a new like record
func (r *likeRepo) Create(ctx context.Context, like *entity.Like) error {
	return r.db.WithContext(ctx).Create(like).Error
}

// GetCount returns the total like count for a given slug
func (r *likeRepo) GetCount(ctx context.Context, slug string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Like{}).Where("slug = ?", slug).Count(&count).Error
	return count, err
}

// ExistsBySlugAndIP checks if a like already exists for the given slug and IP
func (r *likeRepo) ExistsBySlugAndIP(ctx context.Context, slug, ip string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Like{}).
		Where("slug = ? AND ip = ?", slug, ip).
		Count(&count).Error
	return count > 0, err
}
