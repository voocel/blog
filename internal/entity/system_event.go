package entity

import "time"

// EventType defines the type of system event
type EventType string

const (
	EventTypeAudit     EventType = "audit"     // Audit logs for compliance
	EventTypeOperation EventType = "operation" // Operation logs (CRUD operations)
	EventTypeSecurity  EventType = "security"  // Security-related events (login attempts, auth failures)
	EventTypeSystem    EventType = "system"    // System events (startup, shutdown, errors)
	EventTypeBusiness  EventType = "business"  // Business logic events
)

// EventCategory defines the category of event
type EventCategory string

const (
	// Admin-related categories
	CategoryAdminOperation EventCategory = "admin_operation"
	CategoryAdminAuth      EventCategory = "admin_auth"

	// User-related categories
	CategoryUserAction EventCategory = "user_action"
	CategoryUserAuth   EventCategory = "user_auth"

	// System-related categories
	CategorySystemError   EventCategory = "system_error"
	CategorySystemStartup EventCategory = "system_startup"

	// Security-related categories
	CategorySecurityThreat  EventCategory = "security_threat"
	CategorySecurityAttempt EventCategory = "security_attempt"
)

// Severity defines the severity level of event
type Severity string

const (
	SeverityInfo     Severity = "info"
	SeverityWarning  Severity = "warning"
	SeverityError    Severity = "error"
	SeverityCritical Severity = "critical"
)

// SystemEvent records all important system events for monitoring and auditing
type SystemEvent struct {
	ID            int64         `json:"id" gorm:"primaryKey;autoIncrement"`
	RequestID     string        `json:"request_id" gorm:"type:varchar(36);index"`      // Trace request across services
	EventType     EventType     `json:"event_type" gorm:"type:varchar(20);index"`      // audit, operation, security, system, business
	EventCategory EventCategory `json:"event_category" gorm:"type:varchar(50);index"`  // admin_operation, user_action, etc.
	Severity      Severity      `json:"severity" gorm:"type:varchar(20);index"`        // info, warning, error, critical
	UserID        int64         `json:"user_id" gorm:"index"`                          // User ID (if applicable)
	Username      string        `json:"username" gorm:"type:varchar(50)"`              // Username (if applicable)
	Action        string        `json:"action" gorm:"type:varchar(50);index"`          // e.g., "CREATE_POST", "DELETE_USER", "LOGIN_FAILED"
	Resource      string        `json:"resource" gorm:"type:varchar(50)"`              // e.g., "posts", "users", "system"
	ResourceID    int64         `json:"resource_id"`                                   // ID of affected resource
	Method        string        `json:"method" gorm:"type:varchar(10)"`                // HTTP method: POST, PUT, DELETE (for API operations)
	Path          string        `json:"path" gorm:"type:varchar(255)"`                 // Request path (for API operations)
	IP            string        `json:"ip" gorm:"type:varchar(45)"`                    // Client IP address
	UserAgent     string        `json:"user_agent" gorm:"type:varchar(255)"`           // Client user agent
	Status        int           `json:"status"`                                        // HTTP response status code or custom status
	Message       string        `json:"message" gorm:"type:text"`                      // Human-readable message
	ErrorMsg      string        `json:"error_msg" gorm:"type:text"`                    // Error message if failed
	Metadata      string        `json:"metadata" gorm:"type:text"`                     // Additional JSON metadata (flexible for future use)
	Duration      int64         `json:"duration"`                                      // Operation duration in milliseconds
	CreatedAt     time.Time     `json:"created_at" gorm:"autoCreateTime;index"`
}

// TableName specifies the table name for GORM
func (SystemEvent) TableName() string {
	return "system_events"
}

// CreateEventRequest for creating system events
type CreateEventRequest struct {
	RequestID     string
	EventType     EventType
	EventCategory EventCategory
	Severity      Severity
	UserID        int64
	Username      string
	Action        string
	Resource      string
	ResourceID    int64
	Method        string
	Path          string
	IP            string
	UserAgent     string
	Status        int
	Message       string
	ErrorMsg      string
	Metadata      string
	Duration      int64
}

// Helper functions for creating common event types

// NewAuditEvent creates an audit event (for compliance tracking)
func NewAuditEvent(req CreateEventRequest) *SystemEvent {
	req.EventType = EventTypeAudit
	req.Severity = SeverityInfo
	return newEvent(req)
}

// NewOperationEvent creates an operation event (for CRUD operations)
func NewOperationEvent(req CreateEventRequest) *SystemEvent {
	req.EventType = EventTypeOperation
	if req.Severity == "" {
		req.Severity = SeverityInfo
	}
	return newEvent(req)
}

// NewSecurityEvent creates a security event (for auth failures, threats)
func NewSecurityEvent(req CreateEventRequest) *SystemEvent {
	req.EventType = EventTypeSecurity
	if req.Severity == "" {
		req.Severity = SeverityWarning
	}
	return newEvent(req)
}

// NewSystemEvent creates a system event (for system errors, startup)
func NewSystemEvent(req CreateEventRequest) *SystemEvent {
	req.EventType = EventTypeSystem
	if req.Severity == "" {
		req.Severity = SeverityInfo
	}
	return newEvent(req)
}

// newEvent is the internal constructor
func newEvent(req CreateEventRequest) *SystemEvent {
	return &SystemEvent{
		RequestID:     req.RequestID,
		EventType:     req.EventType,
		EventCategory: req.EventCategory,
		Severity:      req.Severity,
		UserID:        req.UserID,
		Username:      req.Username,
		Action:        req.Action,
		Resource:      req.Resource,
		ResourceID:    req.ResourceID,
		Method:        req.Method,
		Path:          req.Path,
		IP:            req.IP,
		UserAgent:     req.UserAgent,
		Status:        req.Status,
		Message:       req.Message,
		ErrorMsg:      req.ErrorMsg,
		Metadata:      req.Metadata,
		Duration:      req.Duration,
	}
}
