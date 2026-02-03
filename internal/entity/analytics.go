package entity

import (
	"time"
)

type Analytics struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	PagePath  string    `gorm:"type:varchar(500);not null;index" json:"pagePath"`
	PostID    *int64    `gorm:"index" json:"postId,omitempty"`
	PostTitle string    `gorm:"type:varchar(255)" json:"postTitle,omitempty"`
	IP        string    `gorm:"type:varchar(45);index" json:"ip"`
	Location  string    `gorm:"type:varchar(200)" json:"location"`
	Timestamp int64     `gorm:"type:bigint;not null;index" json:"timestamp"`
	UserAgent string    `gorm:"type:text" json:"userAgent"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
}

type AnalyticsResponse struct {
	ID        int64  `json:"id"`
	PagePath  string `json:"pagePath"`
	PostID    *int64 `json:"postId,omitempty"`
	PostTitle string `json:"postTitle,omitempty"`
	IP        string `json:"ip"`
	Location  string `json:"location"`
	Timestamp int64  `json:"timestamp"`
	UserAgent string `json:"userAgent"`
}

type LogVisitRequest struct {
	PagePath  string `json:"pagePath" binding:"required"`
	PostID    int64  `json:"postId,omitempty"`
	PostTitle string `json:"postTitle,omitempty"`
}

type DashboardOverviewResponse struct {
	Counts       DashboardCounts       `json:"counts"`
	RecentPosts  []PostResponse        `json:"recentPosts"`
	SystemStatus DashboardSystemStatus `json:"systemStatus"`
}

type DashboardCounts struct {
	Posts      int64 `json:"posts"`      // Total posts including drafts
	Categories int64 `json:"categories"` // Total categories
	Tags       int64 `json:"tags"`       // Total tags
	Files      int64 `json:"files"`      // Total uploaded files
}

type DashboardSystemStatus struct {
	StorageUsage int `json:"storageUsage"` // Storage usage percentage 0-100
	AIQuota      int `json:"aiQuota"`      // AI token usage percentage 0-100
}

func (Analytics) TableName() string {
	return "analytics"
}
