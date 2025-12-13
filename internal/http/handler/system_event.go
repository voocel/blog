package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SystemEventHandler struct {
	eventUseCase *usecase.SystemEventUseCase
}

func NewSystemEventHandler(eventUseCase *usecase.SystemEventUseCase) *SystemEventHandler {
	return &SystemEventHandler{eventUseCase: eventUseCase}
}

// ListEvents - GET /admin/events
func (h *SystemEventHandler) ListEvents(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	userID := c.Query("user_id")
	action := c.Query("action")
	resource := c.Query("resource")
	requestID := c.Query("request_id")
	eventType := c.Query("event_type")         // audit, operation, security, system, business
	eventCategory := c.Query("event_category") // admin_operation, user_action, etc.
	severity := c.Query("severity")            // info, warning, error, critical

	filters := make(map[string]interface{})
	if userID != "" {
		filters["user_id"] = userID
	}
	if action != "" {
		filters["action"] = action
	}
	if resource != "" {
		filters["resource"] = resource
	}
	if requestID != "" {
		filters["request_id"] = requestID
	}
	if eventType != "" {
		filters["event_type"] = eventType
	}
	if eventCategory != "" {
		filters["event_category"] = eventCategory
	}
	if severity != "" {
		filters["severity"] = severity
	}

	result, err := h.eventUseCase.List(c.Request.Context(), filters, page, limit)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetUserEvents - GET /admin/events/user/:id
func (h *SystemEventHandler) GetUserEvents(c *gin.Context) {
	userID := c.Param("id")
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit <= 0 {
		limit = 50
	}

	events, err := h.eventUseCase.GetByUserID(c.Request.Context(), userID, limit)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"events":  events,
		"total":   len(events),
	})
}

// GetRequestTrace - GET /admin/events/trace/:request_id
func (h *SystemEventHandler) GetRequestTrace(c *gin.Context) {
	requestID := c.Param("request_id")

	events, err := h.eventUseCase.GetByRequestID(c.Request.Context(), requestID)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"request_id": requestID,
		"events":     events,
		"total":      len(events),
	})
}

// GetAuditLogs - GET /admin/events/audit (convenience endpoint for audit logs)
func (h *SystemEventHandler) GetAuditLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	result, err := h.eventUseCase.GetAuditLogs(c.Request.Context(), page, limit)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetSecurityEvents - GET /admin/events/security (convenience endpoint for security events)
func (h *SystemEventHandler) GetSecurityEvents(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	result, err := h.eventUseCase.GetSecurityEvents(c.Request.Context(), page, limit)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetSystemErrors - GET /admin/events/errors (convenience endpoint for system errors)
func (h *SystemEventHandler) GetSystemErrors(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit <= 0 {
		limit = 100
	}

	events, err := h.eventUseCase.GetSystemErrors(c.Request.Context(), limit)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"events": events,
		"total":  len(events),
	})
}

// GetEventsByType - GET /admin/events/type/:event_type
func (h *SystemEventHandler) GetEventsByType(c *gin.Context) {
	eventTypeStr := c.Param("event_type")
	eventType := entity.EventType(eventTypeStr)
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit <= 0 {
		limit = 100
	}

	events, err := h.eventUseCase.GetByEventType(c.Request.Context(), eventType, limit)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"event_type": eventType,
		"events":     events,
		"total":      len(events),
	})
}
