package repo

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) usecase.CommentRepo {
	return &commentRepo{db: db}
}

func (r *commentRepo) Create(ctx context.Context, comment *entity.Comment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

func (r *commentRepo) GetByID(ctx context.Context, id string) (*entity.Comment, error) {
	var c entity.Comment
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&c).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}
	return &c, nil
}

func (r *commentRepo) ListTopLevel(ctx context.Context, postID string, page, limit int, order string) ([]entity.Comment, int64, error) {
	var (
		comments []entity.Comment
		total    int64
	)

	q := r.db.WithContext(ctx).Model(&entity.Comment{}).
		Where("post_id = ? AND parent_id IS NULL", postID)

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		q = q.Offset(offset).Limit(limit)
	}

	q = q.Order("created_at " + normalizeOrder(order))

	if err := q.Find(&comments).Error; err != nil {
		return nil, 0, err
	}
	return comments, total, nil
}

func (r *commentRepo) ListReplies(ctx context.Context, postID string, parentIDs []string, order string) ([]entity.Comment, error) {
	if len(parentIDs) == 0 {
		return []entity.Comment{}, nil
	}

	var comments []entity.Comment
	err := r.db.WithContext(ctx).
		Where("post_id = ? AND parent_id IN ?", postID, parentIDs).
		Order("created_at " + normalizeOrder(order)).
		Find(&comments).Error
	return comments, err
}

func normalizeOrder(order string) string {
	switch strings.ToLower(order) {
	case "asc":
		return "ASC"
	default:
		return "DESC"
	}
}
