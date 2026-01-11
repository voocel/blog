package entity

import "time"

// Like represents a like record for any content (homepage, posts, etc.)
type Like struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Slug      string    `gorm:"size:100;index" json:"slug"` // e.g., "home", "post-123"
	IP        string    `gorm:"size:50" json:"ip"`
	UserAgent string    `gorm:"size:255" json:"user_agent"`
	CreatedAt time.Time `json:"created_at"`
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
