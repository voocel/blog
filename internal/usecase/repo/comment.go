package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

func (c CommentRepo) AddCommentRepo(ctx context.Context, comment *entity.Comment) error {
	//TODO implement me
	panic("implement me")
}

func (c CommentRepo) GetCommentByIdRepo(ctx context.Context, id int64) (*entity.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c CommentRepo) GetCommentsByArticleIdRepo(ctx context.Context, aid int64) ([]*entity.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c CommentRepo) GetCommentsRepo(ctx context.Context) ([]*entity.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c CommentRepo) DeleteCommentRepo(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
