package usecase

import (
	"blog/internal/entity"
	"context"
)

type CommentUseCase struct {
	repo CommentRepo
}

func NewCommentUseCase(repo CommentRepo) *CommentUseCase {
	return &CommentUseCase{repo: repo}
}

func (c *CommentUseCase) AddComment(ctx context.Context, comment *entity.Comment) error {
	return c.repo.AddCommentRepo(ctx, comment)
}

func (c *CommentUseCase) GetCommentById(ctx context.Context, cid int64) (*entity.Comment, error) {
	return c.repo.GetCommentByIdRepo(ctx, cid)
}

func (c *CommentUseCase) GetComments(ctx context.Context) ([]*entity.Comment, error) {
	return c.repo.GetCommentsRepo(ctx)
}

func (c *CommentUseCase) GetCommentsByArticleId(ctx context.Context, aid int64) ([]*entity.Comment, error) {
	return c.repo.GetCommentsByArticleIdRepo(ctx, aid)
}

func (c *CommentUseCase) DeleteComment(ctx context.Context, cid int64) error {
	return c.repo.DeleteCommentRepo(ctx, cid)
}

func (c *CommentUseCase) GetCommentsWithPagination(ctx context.Context, page, pageSize int, articleId, discussionId *int64) ([]*entity.Comment, int64, error) {
	return c.repo.GetCommentsWithPaginationRepo(ctx, page, pageSize, articleId, discussionId)
}

func (c *CommentUseCase) UpdateComment(ctx context.Context, comment *entity.Comment) error {
	return c.repo.UpdateCommentRepo(ctx, comment)
}
