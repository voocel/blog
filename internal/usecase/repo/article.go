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

// List supports multiple filter conditions and optional pagination
func (r *postRepo) List(ctx context.Context, filters map[string]interface{}, page, limit int) ([]entity.Post, int64, error) {
	var posts []entity.Post
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Post{})

	// Apply filter conditions
	if categoryID, ok := filters["categoryId"].(string); ok && categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if search, ok := filters["search"].(string); ok && search != "" {
		query = query.Where("title LIKE ? OR excerpt LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	// Filter by date (for scheduled publishing: only show posts with date <= current date)
	if beforeDate, ok := filters["beforeDate"].(string); ok && beforeDate != "" {
		query = query.Where("date <= ?", beforeDate)
	}

	// Get total count
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Paginated query (if page and limit provided)
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	// Order by date descending
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
	// First delete associated tags
	r.RemoveTags(ctx, id)
	// Then delete the post
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Post{}).Error
}

func (r *postRepo) IncrementViews(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&entity.Post{}).
		Where("id = ?", id).
		UpdateColumn("views", gorm.Expr("views + 1")).Error
}

// AddTags adds post-tag associations
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

// RemoveTags removes all tag associations for a post
func (r *postRepo) RemoveTags(ctx context.Context, postID string) error {
	return r.db.WithContext(ctx).Where("post_id = ?", postID).Delete(&entity.PostTag{}).Error
}

// GetTagIDs gets tag ID list for a post
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

// Count counts total number of posts (including drafts)
func (r *postRepo) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Post{}).Count(&count).Error
	return count, err
}

// GetRecent gets recent posts
func (r *postRepo) GetRecent(ctx context.Context, limit int) ([]entity.Post, error) {
	var posts []entity.Post
	err := r.db.WithContext(ctx).
		Model(&entity.Post{}).
		Order("date DESC").
		Limit(limit).
		Find(&posts).Error
	return posts, err
}
