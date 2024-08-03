package usecase

import (
	"blog/internal/entity"
	"context"
)

type StarUseCase struct {
	repo StarRepo
}

func NewStarUseCase(repo StarRepo) *StarUseCase {
	return &StarUseCase{repo: repo}
}

func (c *StarUseCase) AddStar(ctx context.Context, uid, aid int64) error {
	return c.repo.AddStarRepo(ctx, uid, aid)
}

func (c *StarUseCase) DeleteStar(ctx context.Context, uid, aid int64) error {
	return c.repo.DeleteStarRepo(ctx, uid, aid)
}

func (c *StarUseCase) GetStarsByUidRepo(ctx context.Context, uid int64) ([]*entity.Star, error) {
	return c.repo.GetStarsByUidRepo(ctx, uid)
}

func (c *StarUseCase) GetStarsByArticleId(ctx context.Context, aid int64) ([]*entity.Star, error) {
	return c.repo.GetStarsByArticleIdRepo(ctx, aid)
}
