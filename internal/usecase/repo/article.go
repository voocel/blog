package repo

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"context"
	"errors"
	"time"

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

func (r *postRepo) CreateWithTags(ctx context.Context, post *entity.Post, tagIDs []string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(post).Error; err != nil {
			return err
		}
		if len(tagIDs) == 0 {
			return nil
		}
		postTags := make([]entity.PostTag, 0, len(tagIDs))
		for _, tagID := range tagIDs {
			postTags = append(postTags, entity.PostTag{
				PostID: post.ID,
				TagID:  tagID,
			})
		}
		return tx.Create(&postTags).Error
	})
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

func (r *postRepo) GetBySlug(ctx context.Context, slug string) (*entity.Post, error) {
	var post entity.Post
	err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &post, nil
}

func (r *postRepo) SlugExists(ctx context.Context, slug string, excludeID string) (bool, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&entity.Post{}).Where("slug = ?", slug)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *postRepo) GetByIDs(ctx context.Context, ids []string) ([]entity.Post, error) {
	if len(ids) == 0 {
		return []entity.Post{}, nil
	}
	var posts []entity.Post
	if err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
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
	if beforePublishAt, ok := filters["beforePublishAt"].(time.Time); ok && !beforePublishAt.IsZero() {
		query = query.Where("publish_at <= ?", beforePublishAt)
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

	// Order by publish time descending
	err = query.Order("publish_at DESC").Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (r *postRepo) Update(ctx context.Context, post *entity.Post) error {
	return r.db.WithContext(ctx).Save(post).Error
}

func (r *postRepo) UpdateWithTags(ctx context.Context, post *entity.Post, tagIDs []string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(post).Error; err != nil {
			return err
		}
		// Replace all tag associations
		if err := tx.Where("post_id = ?", post.ID).Delete(&entity.PostTag{}).Error; err != nil {
			return err
		}
		if len(tagIDs) == 0 {
			return nil
		}
		postTags := make([]entity.PostTag, 0, len(tagIDs))
		for _, tagID := range tagIDs {
			postTags = append(postTags, entity.PostTag{
				PostID: post.ID,
				TagID:  tagID,
			})
		}
		return tx.Create(&postTags).Error
	})
}

func (r *postRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("post_id = ?", id).Delete(&entity.Comment{}).Error; err != nil {
			return err
		}
		if err := tx.Where("post_id = ?", id).Delete(&entity.PostTag{}).Error; err != nil {
			return err
		}
		return tx.Where("id = ?", id).Delete(&entity.Post{}).Error
	})
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

func (r *postRepo) GetTagIDsByPostIDs(ctx context.Context, postIDs []string) (map[string][]string, error) {
	if len(postIDs) == 0 {
		return map[string][]string{}, nil
	}
	var postTags []entity.PostTag
	if err := r.db.WithContext(ctx).Where("post_id IN ?", postIDs).Find(&postTags).Error; err != nil {
		return nil, err
	}
	out := make(map[string][]string, len(postIDs))
	for _, pt := range postTags {
		out[pt.PostID] = append(out[pt.PostID], pt.TagID)
	}
	return out, nil
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
		Order("publish_at DESC").
		Limit(limit).
		Find(&posts).Error
	return posts, err
}
