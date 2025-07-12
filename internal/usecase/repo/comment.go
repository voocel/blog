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

func (c CommentRepo) GetCommentsWithPaginationRepo(ctx context.Context, page, pageSize int, articleId, discussionId *int64) ([]*entity.Comment, int64, error) {
	var comments []*entity.Comment
	var total int64

	query := c.db.WithContext(ctx).Model(&entity.Comment{})

	// 添加过滤条件
	if articleId != nil {
		query = query.Where("article_id = ?", *articleId)
	}
	if discussionId != nil {
		query = query.Where("discussion_id = ?", *discussionId)
	}

	// 获取总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

func (c CommentRepo) UpdateCommentRepo(ctx context.Context, comment *entity.Comment) error {
	return c.db.WithContext(ctx).Save(comment).Error
}
