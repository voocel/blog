package usecase

import (
	"blog/internal/entity"
	"context"
)

// UserRepo user repository interface
type UserRepo interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}

// PostRepo post repository interface
type PostRepo interface {
	Create(ctx context.Context, post *entity.Post) error
	GetByID(ctx context.Context, id string) (*entity.Post, error)
	List(ctx context.Context, filters map[string]interface{}, page, limit int) ([]entity.Post, int64, error)
	Update(ctx context.Context, post *entity.Post) error
	Delete(ctx context.Context, id string) error
	IncrementViews(ctx context.Context, id string) error

	// Tag associations
	AddTags(ctx context.Context, postID string, tagIDs []string) error
	RemoveTags(ctx context.Context, postID string) error
	GetTagIDs(ctx context.Context, postID string) ([]string, error)

	// Statistics
	Count(ctx context.Context) (int64, error)
	GetRecent(ctx context.Context, limit int) ([]entity.Post, error)
}

// CategoryRepo category repository interface
type CategoryRepo interface {
	Create(ctx context.Context, category *entity.Category) error
	GetByID(ctx context.Context, id string) (*entity.Category, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Category, error)
	List(ctx context.Context) ([]entity.Category, error)
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id string) error
	IncrementCount(ctx context.Context, id string) error
	DecrementCount(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}

// TagRepo tag repository interface
type TagRepo interface {
	Create(ctx context.Context, tag *entity.Tag) error
	GetByID(ctx context.Context, id string) (*entity.Tag, error)
	GetByName(ctx context.Context, name string) (*entity.Tag, error)
	GetByIDs(ctx context.Context, ids []string) ([]entity.Tag, error)
	List(ctx context.Context) ([]entity.Tag, error)
	Update(ctx context.Context, tag *entity.Tag) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}

// MediaRepo media repository interface
type MediaRepo interface {
	Create(ctx context.Context, media *entity.Media) error
	GetByID(ctx context.Context, id string) (*entity.Media, error)
	List(ctx context.Context) ([]entity.Media, error)
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}

// AnalyticsRepo analytics repository interface
type AnalyticsRepo interface {
	Create(ctx context.Context, log *entity.Analytics) error
	GetLogs(ctx context.Context, startDate, endDate string, limit int) ([]entity.Analytics, error)
}
