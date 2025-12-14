package entity

import (
	"time"
)

type Post struct {
	ID      string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Excerpt string `gorm:"type:varchar(500);not null" json:"excerpt"`
	Content string `gorm:"type:text;not null" json:"content"` // Markdown
	Author  string `gorm:"type:varchar(100);not null" json:"author"`
	// PublishAt is the scheduled publish time (supports second-level scheduling).
	// Public APIs only return posts where status=published and publish_at <= now().
	PublishAt  time.Time `gorm:"not null;index" json:"publishAt"`
	CategoryID string    `gorm:"type:uuid;not null;index" json:"categoryId"`
	Cover      string    `gorm:"type:varchar(500);not null" json:"cover"`
	Views      int       `gorm:"type:int;default:0" json:"views"`
	Status     string    `gorm:"type:varchar(20);not null;default:'draft'" json:"status"` // published | draft
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"-"`
}

type PostTag struct {
	ID        string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PostID    string    `gorm:"type:uuid;not null;index" json:"postId"`
	TagID     string    `gorm:"type:uuid;not null;index" json:"tagId"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

type PostResponse struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Excerpt    string    `json:"excerpt"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
	PublishAt  time.Time `json:"publishAt"`
	CategoryID string    `json:"categoryId"`
	Category   string    `json:"category"` // Category name
	ReadTime   string    `json:"readTime"`
	Cover      string    `json:"cover"`
	Tags       []string  `json:"tags"` // Tag names
	Views      int       `json:"views"`
	Status     string    `json:"status"`
}

type CreatePostRequest struct {
	Title      string   `json:"title" binding:"required"`
	Excerpt    string   `json:"excerpt"` // Optional: if empty, backend will derive from content
	Content    string   `json:"content" binding:"required"`
	CategoryID string   `json:"categoryId" binding:"required"`
	Tags       []string `json:"tags"` // Tag IDs
	Cover      string   `json:"cover" binding:"required"`
	Status     string   `json:"status"` // published | draft, default: draft
	// PublishAt should be RFC3339 (e.g. 2025-12-14T16:30:00+08:00).
	// If omitted, defaults to server current time.
	PublishAt string `json:"publishAt"`
}

type UpdatePostRequest struct {
	Title      string   `json:"title,omitempty"`
	Excerpt    string   `json:"excerpt,omitempty"`
	Content    string   `json:"content,omitempty"`
	CategoryID string   `json:"categoryId,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	Cover      string   `json:"cover,omitempty"`
	Status     string   `json:"status,omitempty"`
	PublishAt  string   `json:"publishAt,omitempty"` // RFC3339
}

type PaginatedPostsResponse struct {
	Data       []PostResponse `json:"data"`
	Pagination Pagination     `json:"pagination"`
}

type Pagination struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
}

func (Post) TableName() string {
	return "posts"
}

func (PostTag) TableName() string {
	return "post_tags"
}
