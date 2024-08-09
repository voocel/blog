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
	return c.db.WithContext(ctx).Create(comment).Error
}

func (c CommentRepo) GetCommentByIdRepo(ctx context.Context, id int64) (*entity.Comment, error) {
	var comment = new(entity.Comment)
	return comment, c.db.WithContext(ctx).Where("id = ?", id).First(comment).Error
}

func (c CommentRepo) GetCommentsByArticleIdRepo(ctx context.Context, aid int64) ([]*entity.Comment, error) {
	var comments []*entity.Comment
	return comments, c.db.WithContext(ctx).Where("article_id = ?", aid).Find(&comments).Error
}

func (c CommentRepo) GetCommentsRepo(ctx context.Context) ([]*entity.Comment, error) {
	var comments []*entity.Comment
	return comments, c.db.WithContext(ctx).Find(&comments).Error
}

func (c CommentRepo) DeleteCommentRepo(ctx context.Context, id int64) error {
	return c.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Comment{}).Error
}
