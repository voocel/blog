package entity

import "time"

// Like represents a like record for any content (homepage, posts, etc.)
type Like struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Slug      string    `gorm:"size:100;index" json:"slug"` // e.g., "home", "post-123"
	IP        string    `gorm:"size:50" json:"ip"`
	UserAgent string    `gorm:"size:255" json:"userAgent"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

func (Like) TableName() string {
	return "likes"
}

// LikeCount represents aggregated like count for a slug
type LikeCount struct {
	Slug  string `json:"slug"`
	Count int64  `json:"count"`
}

// LikeResponse is the API response for like operations
type LikeResponse struct {
	Count int64 `json:"count"`
}
