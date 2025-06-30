package entity

import "time"

// ApiResponse 统一API响应格式
type ApiResponse[T any] struct {
	Code      int    `json:"code"`      // 状态码：200成功，其他为错误
	Message   string `json:"message"`   // 响应消息
	Data      T      `json:"data"`      // 响应数据
	Timestamp int64  `json:"timestamp"` // 时间戳
}

// PaginatedResponse 分页响应格式
type PaginatedResponse[T any] struct {
	Items      []T `json:"items"`      // 数据列表
	Total      int `json:"total"`      // 总数量
	Page       int `json:"page"`       // 当前页码
	PageSize   int `json:"pageSize"`   // 每页大小
	TotalPages int `json:"totalPages"` // 总页数
}

// ErrorDetail 详细错误信息
type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Code      int           `json:"code"`
	Message   string        `json:"message"`
	Errors    []ErrorDetail `json:"errors,omitempty"` // 详细错误信息（可选）
	Timestamp int64         `json:"timestamp"`
}

// NewSuccessResponse 创建成功响应
func NewSuccessResponse[T any](data T, message string) *ApiResponse[T] {
	return &ApiResponse[T]{
		Code:      200,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
}

// NewErrorResponse 创建错误响应
func NewErrorResponse(code int, message string, errors ...ErrorDetail) *ErrorResponse {
	return &ErrorResponse{
		Code:      code,
		Message:   message,
		Errors:    errors,
		Timestamp: time.Now().Unix(),
	}
}

// NewPaginatedResponse 创建分页响应
func NewPaginatedResponse[T any](items []T, total, page, pageSize int) *PaginatedResponse[T] {
	totalPages := (total + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}

	return &PaginatedResponse[T]{
		Items:      items,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}

// 友链相关响应结构
type FriendlinkRequest struct {
	Name        string `json:"name" binding:"required"`
	URL         string `json:"url" binding:"required,url"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"oneof=active inactive"`
	SortOrder   int    `json:"sortOrder"`
}

type FriendlinkResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Status      string `json:"status"`
	SortOrder   int    `json:"sortOrder"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// 统计相关响应结构
type StatisticsItem struct {
	Total  int     `json:"total"`
	Growth float64 `json:"growth"`
}

type DashboardStatistics struct {
	Users    StatisticsItem `json:"users"`
	Articles StatisticsItem `json:"articles"`
	Comments StatisticsItem `json:"comments"`
	Visits   StatisticsItem `json:"visits"`
}

type VisitStatistics struct {
	ID           string `json:"id"`
	ArticleID    string `json:"articleId,omitempty"`
	ArticleTitle string `json:"articleTitle,omitempty"`
	IP           string `json:"ip"`
	UserAgent    string `json:"userAgent"`
	Referer      string `json:"referer,omitempty"`
	VisitCount   int    `json:"visitCount"`
	CreatedAt    string `json:"createdAt"`
}

// 系统信息相关响应结构
type DatabaseInfo struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

type PHPInfo struct {
	Version    string   `json:"version"`
	Extensions []string `json:"extensions"`
}

type SystemInfo struct {
	Language  string       `json:"language"`
	Version   string       `json:"version"`
	WebServer string       `json:"webServer"`
	Domain    string       `json:"domain"`
	IP        string       `json:"ip"`
	UserAgent string       `json:"userAgent"`
	Database  DatabaseInfo `json:"database"`
	PHP       *PHPInfo     `json:"php,omitempty"`
}
