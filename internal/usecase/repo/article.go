package repo

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"context"
	"errors"

	"gorm.io/gorm"
)

type postRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) usecase.PostRepo {
	return &postRepo{db: db}
}

func (r *postRepo) Create(ctx context.Context, post *entity.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *postRepo) GetByID(ctx context.Context, id string) (*entity.Post, error) {
	var post entity.Post
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &post, nil
}

// List 支持多种过滤条件和可选分页
func (r *postRepo) List(ctx context.Context, filters map[string]interface{}, page, limit int) ([]entity.Post, int64, error) {
	var posts []entity.Post
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Post{})

	// 应用过滤条件
	if categoryID, ok := filters["categoryId"].(string); ok && categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if search, ok := filters["search"].(string); ok && search != "" {
		query = query.Where("title LIKE ? OR excerpt LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// 获取总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询（如果提供了 page 和 limit）
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	// 按日期倒序排列
	err = query.Order("date DESC").Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (r *postRepo) Update(ctx context.Context, post *entity.Post) error {
	return r.db.WithContext(ctx).Save(post).Error
}

func (r *postRepo) Delete(ctx context.Context, id string) error {
	// 先删除关联的标签
	r.RemoveTags(ctx, id)
	// 再删除文章
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Post{}).Error
}

func (r *postRepo) IncrementViews(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&entity.Post{}).
		Where("id = ?", id).
		UpdateColumn("views", gorm.Expr("views + 1")).Error
}

// AddTags 添加文章标签关联
func (r *postRepo) AddTags(ctx context.Context, postID string, tagIDs []string) error {
	var postTags []entity.PostTag
	for _, tagID := range tagIDs {
		postTags = append(postTags, entity.PostTag{
			PostID: postID,
			TagID:  tagID,
		})
	}
	return r.db.WithContext(ctx).Create(&postTags).Error
}

// RemoveTags 删除文章的所有标签关联
func (r *postRepo) RemoveTags(ctx context.Context, postID string) error {
	return r.db.WithContext(ctx).Where("post_id = ?", postID).Delete(&entity.PostTag{}).Error
}

// GetTagIDs 获取文章的标签 ID 列表
func (r *postRepo) GetTagIDs(ctx context.Context, postID string) ([]string, error) {
	var postTags []entity.PostTag
	err := r.db.WithContext(ctx).Where("post_id = ?", postID).Find(&postTags).Error
	if err != nil {
		return nil, err
	}

	tagIDs := make([]string, len(postTags))
	for i, pt := range postTags {
		tagIDs[i] = pt.TagID
	}
	return tagIDs, nil
}
