package usecase

import (
	"blog/internal/entity"
	"context"
)

type SystemEventUseCase struct {
	eventRepo SystemEventRepo
}

func NewSystemEventUseCase(eventRepo SystemEventRepo) *SystemEventUseCase {
	return &SystemEventUseCase{eventRepo: eventRepo}
}

// List retrieves paginated system events with optional filters
func (uc *SystemEventUseCase) List(ctx context.Context, filters map[string]interface{}, page, limit int) (map[string]interface{}, error) {
	events, total, err := uc.eventRepo.List(ctx, filters, page, limit)
	if err != nil {
		return nil, err
	}

	if limit <= 0 {
		limit = 20
	}
	if page <= 0 {
		page = 1
	}

	return map[string]interface{}{
		"events": events,
		"total":  total,
		"page":   page,
		"limit":  limit,
		"pages":  (total + int64(limit) - 1) / int64(limit),
	}, nil
}

// GetByRequestID retrieves all events for a specific request ID (distributed tracing)
func (uc *SystemEventUseCase) GetByRequestID(ctx context.Context, requestID string) ([]entity.SystemEvent, error) {
	return uc.eventRepo.GetByRequestID(ctx, requestID)
}

// GetByUserID retrieves recent events for a specific user
func (uc *SystemEventUseCase) GetByUserID(ctx context.Context, userID int64, limit int) ([]entity.SystemEvent, error) {
	return uc.eventRepo.GetByUserID(ctx, userID, limit)
}

// GetByEventType retrieves events by type (audit, operation, security, system, business)
func (uc *SystemEventUseCase) GetByEventType(ctx context.Context, eventType entity.EventType, limit int) ([]entity.SystemEvent, error) {
	return uc.eventRepo.GetByEventType(ctx, eventType, limit)
}

// GetAuditLogs is a convenience method for retrieving audit logs
func (uc *SystemEventUseCase) GetAuditLogs(ctx context.Context, page, limit int) (map[string]interface{}, error) {
	filters := map[string]interface{}{
		"event_type": entity.EventTypeAudit,
	}
	return uc.List(ctx, filters, page, limit)
}

// GetSecurityEvents is a convenience method for retrieving security events
func (uc *SystemEventUseCase) GetSecurityEvents(ctx context.Context, page, limit int) (map[string]interface{}, error) {
	filters := map[string]interface{}{
		"event_type": entity.EventTypeSecurity,
	}
	return uc.List(ctx, filters, page, limit)
}

// GetSystemErrors is a convenience method for retrieving system errors
func (uc *SystemEventUseCase) GetSystemErrors(ctx context.Context, limit int) ([]entity.SystemEvent, error) {
	filters := map[string]interface{}{
		"event_type": entity.EventTypeSystem,
		"severity":   entity.SeverityError,
	}
	events, _, err := uc.eventRepo.List(ctx, filters, 1, limit)
	return events, err
}
