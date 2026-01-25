package usecase

import (
	"blog/internal/entity"
	"context"
	"errors"
)

var ErrAlreadyLiked = errors.New("already liked from this IP")

type LikeUseCase struct {
	likeRepo LikeRepo
}

func NewLikeUseCase(likeRepo LikeRepo) *LikeUseCase {
	return &LikeUseCase{likeRepo: likeRepo}
}

// GetCount returns the like count for a given slug
func (uc *LikeUseCase) GetCount(ctx context.Context, slug string) (int64, error) {
	return uc.likeRepo.GetCount(ctx, slug)
}

// Like adds a new like and returns the updated count
// Returns ErrAlreadyLiked if the IP has already liked this slug
func (uc *LikeUseCase) Like(ctx context.Context, slug, ip, userAgent string) (int64, error) {
	// Check if this IP has already liked
	exists, err := uc.likeRepo.ExistsBySlugAndIP(ctx, slug, ip)
	if err != nil {
		return 0, err
	}
	if exists {
		// Return current count without error for idempotent behavior
		return uc.likeRepo.GetCount(ctx, slug)
	}

	like := &entity.Like{
		Slug:      slug,
		IP:        ip,
		UserAgent: userAgent,
	}

	if err := uc.likeRepo.Create(ctx, like); err != nil {
		return 0, err
	}

	return uc.likeRepo.GetCount(ctx, slug)
}
