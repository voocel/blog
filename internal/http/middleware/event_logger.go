package middleware

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// EventLogger logs important system events (admin operations, auth attempts, etc.)
func EventLogger(eventRepo usecase.SystemEventRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only log important operations
		if !shouldLogEvent(c) {
			c.Next()
			return
		}

		// Capture request body for important operations
		var metadata string
		if c.Request.Method != "GET" && c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			metadata = string(bodyBytes)
			if len(metadata) > 5000 {
				metadata = metadata[:5000] + "... (truncated)"
			}
		}

		userID, _ := c.Get("user_id")
		username, _ := c.Get("username")
		requestID, _ := c.Get("request_id")

		startTime := time.Now()

		// Create response writer wrapper to capture status code
		writer := &responseWriter{
			ResponseWriter: c.Writer,
			statusCode:     200, // Default status
		}
		c.Writer = writer
		c.Next()

		duration := time.Since(startTime).Milliseconds()
		var errorMsg string
		if len(c.Errors) > 0 {
			errorMsg = c.Errors.String()
		}

		// Determine event details from path and method
		eventType, eventCategory, action, resource, resourceID := parseEventInfo(c)
		severity := determineSeverity(writer.statusCode, c)
		message := generateEventMessage(action, resource, getStringValue(username), writer.statusCode)

		// Create system event
		event := &entity.SystemEvent{
			ID:            uuid.New().String(),
			RequestID:     getStringValue(requestID),
			EventType:     eventType,
			EventCategory: eventCategory,
			Severity:      severity,
			UserID:        getStringValue(userID),
			Username:      getStringValue(username),
			Action:        action,
			Resource:      resource,
			ResourceID:    resourceID,
			Method:        c.Request.Method,
			Path:          c.Request.URL.Path,
			IP:            c.ClientIP(),
			UserAgent:     c.Request.UserAgent(),
			Status:        writer.statusCode,
			Message:       message,
			ErrorMsg:      errorMsg,
			Metadata:      metadata,
			Duration:      duration,
		}

		// Async logging to avoid blocking response
		go func() {
			_ = eventRepo.Create(c.Request.Context(), event)
		}()
	}
}

// responseWriter wraps gin.ResponseWriter to capture status code
type responseWriter struct {
	gin.ResponseWriter
	statusCode int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// shouldLogEvent determines if the request should be logged
func shouldLogEvent(c *gin.Context) bool {
	path := c.Request.URL.Path

	// Log all admin operations
	if strings.HasPrefix(path, "/api/v1/admin") {
		return true
	}

	// Log authentication operations
	if strings.HasPrefix(path, "/api/v1/auth/") {
		return true
	}

	// Log profile updates
	if strings.HasPrefix(path, "/api/v1/users/profile") && c.Request.Method != "GET" {
		return true
	}

	return false
}

// parseEventInfo extracts event details from request
func parseEventInfo(c *gin.Context) (entity.EventType, entity.EventCategory, string, string, string) {
	path := c.Request.URL.Path
	method := c.Request.Method

	// Default values
	eventType := entity.EventTypeOperation
	eventCategory := entity.CategoryAdminOperation
	var resource, resourceID string

	// Determine event type and category
	if strings.HasPrefix(path, "/api/v1/auth/") {
		eventType = entity.EventTypeSecurity
		if strings.Contains(path, "/admin/") {
			eventCategory = entity.CategoryAdminAuth
		} else {
			eventCategory = entity.CategoryUserAuth
		}
		resource = "auth"
	} else if strings.HasPrefix(path, "/api/v1/admin/") {
		eventType = entity.EventTypeAudit
		eventCategory = entity.CategoryAdminOperation
		// Extract resource from path
		parts := strings.Split(strings.Trim(path, "/"), "/")
		if len(parts) >= 4 {
			resource = parts[3]
			// Check if last part is an ID
			if len(parts) >= 5 && (len(parts[4]) == 36 || isNumeric(parts[4])) {
				resourceID = parts[4]
			}
		}
	} else {
		eventCategory = entity.CategoryUserAction
		resource = "user"
	}

	// Determine action from HTTP method and path
	action := generateAction(method, path, resource, resourceID)

	return eventType, eventCategory, action, resource, resourceID
}

// generateAction creates action string from method and context
func generateAction(method, path, resource, resourceID string) string {
	// Special handling for auth endpoints
	if strings.Contains(path, "/login") {
		if strings.Contains(path, "failed") {
			return "LOGIN_FAILED"
		}
		return "LOGIN"
	}
	if strings.Contains(path, "/register") {
		return "REGISTER"
	}
	if strings.Contains(path, "/refresh") {
		return "REFRESH_TOKEN"
	}
	if strings.Contains(path, "/logout") {
		return "LOGOUT"
	}

	// Special handling for upload endpoint
	if strings.Contains(path, "/upload") {
		return "UPLOAD_FILE"
	}

	// Standard CRUD actions
	switch method {
	case "POST":
		return "CREATE_" + strings.ToUpper(resource)
	case "PUT", "PATCH":
		return "UPDATE_" + strings.ToUpper(resource)
	case "DELETE":
		return "DELETE_" + strings.ToUpper(resource)
	case "GET":
		if resourceID != "" {
			return "VIEW_" + strings.ToUpper(resource)
		}
		return "LIST_" + strings.ToUpper(resource)
	default:
		return method + "_" + strings.ToUpper(resource)
	}
}

// determineSeverity determines event severity based on status code and context
func determineSeverity(statusCode int, c *gin.Context) entity.Severity {
	path := c.Request.URL.Path

	// Critical: Failed authentication attempts
	if strings.Contains(path, "/auth/login") && statusCode >= 400 {
		return entity.SeverityCritical
	}

	// Error: 5xx server errors
	if statusCode >= 500 {
		return entity.SeverityError
	}

	// Warning: 4xx client errors
	if statusCode >= 400 {
		return entity.SeverityWarning
	}

	// Info: successful operations
	return entity.SeverityInfo
}

// generateEventMessage creates a human-readable message
func generateEventMessage(action, resource, username string, statusCode int) string {
	user := username
	if user == "" {
		user = "Anonymous"
	}

	status := "successfully"
	if statusCode >= 400 {
		status = "failed to"
	}

	return user + " " + status + " " + strings.ToLower(strings.ReplaceAll(action, "_", " "))
}

// getStringValue safely converts interface{} to string
func getStringValue(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

// isNumeric checks if string is numeric
func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(s) > 0
}
