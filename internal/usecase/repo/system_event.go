package repo

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"context"

	"gorm.io/gorm"
)

type systemEventRepo struct {
	db *gorm.DB
}

func NewSystemEventRepo(db *gorm.DB) usecase.SystemEventRepo {
	return &systemEventRepo{db: db}
}

func (r *systemEventRepo) Create(ctx context.Context, event *entity.SystemEvent) error {
	return r.db.WithContext(ctx).Create(event).Error
}

func (r *systemEventRepo) List(ctx context.Context, filters map[string]interface{}, page, limit int) ([]entity.SystemEvent, int64, error) {
	var events []entity.SystemEvent
	var total int64

	if limit <= 0 {
		limit = 20
	}
	if page <= 0 {
		page = 1
	}

	query := r.db.WithContext(ctx).Model(&entity.SystemEvent{})

	// Apply filters
	if userID, ok := filters["user_id"].(int64); ok && userID != 0 {
		query = query.Where("user_id = ?", userID)
	}
	if action, ok := filters["action"].(string); ok && action != "" {
		query = query.Where("action = ?", action)
	}
	if resource, ok := filters["resource"].(string); ok && resource != "" {
		query = query.Where("resource = ?", resource)
	}
	if requestID, ok := filters["request_id"].(string); ok && requestID != "" {
		query = query.Where("request_id = ?", requestID)
	}
	if eventType, ok := filters["event_type"].(string); ok && eventType != "" {
		query = query.Where("event_type = ?", eventType)
	}
	if eventCategory, ok := filters["event_category"].(string); ok && eventCategory != "" {
		query = query.Where("event_category = ?", eventCategory)
	}
	if severity, ok := filters["severity"].(string); ok && severity != "" {
		query = query.Where("severity = ?", severity)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, total, nil
}

func (r *systemEventRepo) GetByRequestID(ctx context.Context, requestID string) ([]entity.SystemEvent, error) {
	var events []entity.SystemEvent
	err := r.db.WithContext(ctx).
		Where("request_id = ?", requestID).
		Order("created_at ASC").
		Find(&events).Error
	return events, err
}

func (r *systemEventRepo) GetByUserID(ctx context.Context, userID int64, limit int) ([]entity.SystemEvent, error) {
	if limit <= 0 {
		limit = 50
	}

	var events []entity.SystemEvent
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&events).Error
	return events, err
}

func (r *systemEventRepo) GetByEventType(ctx context.Context, eventType entity.EventType, limit int) ([]entity.SystemEvent, error) {
	if limit <= 0 {
		limit = 100
	}

	var events []entity.SystemEvent
	err := r.db.WithContext(ctx).
		Where("event_type = ?", eventType).
		Order("created_at DESC").
		Limit(limit).
		Find(&events).Error
	return events, err
}
