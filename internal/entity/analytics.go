package entity

import (
	"time"
)

// Analytics 访问统计模型
type Analytics struct {
	ID        string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PagePath  string    `gorm:"type:varchar(500);not null;index" json:"pagePath"`
	PostID    *string   `gorm:"type:uuid;index" json:"postId,omitempty"` // 允许为空，防止空串触发 UUID 解析错误
	PostTitle string    `gorm:"type:varchar(255)" json:"postTitle,omitempty"`
	IP        string    `gorm:"type:varchar(45);index" json:"ip"`
	Location  string    `gorm:"type:varchar(200)" json:"location"`
	Timestamp int64     `gorm:"type:bigint;not null;index" json:"timestamp"`
	UserAgent string    `gorm:"type:text" json:"userAgent"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
}

// AnalyticsResponse 访问统计响应
type AnalyticsResponse struct {
	ID        string  `json:"id"`
	PagePath  string  `json:"pagePath"`
	PostID    *string `json:"postId,omitempty"`
	PostTitle string  `json:"postTitle,omitempty"`
	IP        string  `json:"ip"`
	Location  string  `json:"location"`
	Timestamp int64   `json:"timestamp"`
	UserAgent string  `json:"userAgent"`
}

// LogVisitRequest 记录访问请求
type LogVisitRequest struct {
	PagePath  string `json:"pagePath" binding:"required"`
	PostID    string `json:"postId,omitempty"`
	PostTitle string `json:"postTitle,omitempty"`
}

func (Analytics) TableName() string {
	return "analytics"
}
