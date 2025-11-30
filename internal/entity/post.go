package entity

import (
	"time"
)

// Post 博客文章模型
type Post struct {
	ID         string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title      string    `gorm:"type:varchar(255);not null" json:"title"`
	Excerpt    string    `gorm:"type:varchar(500);not null" json:"excerpt"`
	Content    string    `gorm:"type:text;not null" json:"content"` // Markdown
	Author     string    `gorm:"type:varchar(100);not null" json:"author"`
	Date       string    `gorm:"type:varchar(20);not null" json:"date"` // ISO 8601 format: '2024-05-15'
	CategoryID string    `gorm:"type:uuid;not null;index" json:"categoryId"`
	ImageUrl   string    `gorm:"type:varchar(500);not null" json:"imageUrl"`
	Views      int       `gorm:"type:int;default:0" json:"views"`
	Status     string    `gorm:"type:varchar(20);not null;default:'draft'" json:"status"` // published | draft
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"-"`
}

// PostTag 文章标签关联表
type PostTag struct {
	ID        string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PostID    string    `gorm:"type:uuid;not null;index" json:"postId"`
	TagID     string    `gorm:"type:uuid;not null;index" json:"tagId"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

// PostResponse 文章响应（包含关联信息）
type PostResponse struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Excerpt    string   `json:"excerpt"`
	Content    string   `json:"content"`
	Author     string   `json:"author"`
	Date       string   `json:"date"`
	CategoryID string   `json:"categoryId"`
	Category   string   `json:"category"` // Category name
	ReadTime   string   `json:"readTime"`
	ImageUrl   string   `json:"imageUrl"`
	Tags       []string `json:"tags"` // Tag names
	Views      int      `json:"views"`
	Status     string   `json:"status"`
}

// CreatePostRequest 创建文章请求
type CreatePostRequest struct {
	Title      string   `json:"title" binding:"required"`
	Excerpt    string   `json:"excerpt" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	CategoryID string   `json:"categoryId" binding:"required"`
	Tags       []string `json:"tags"` // Tag IDs
	ImageUrl   string   `json:"imageUrl" binding:"required"`
	Status     string   `json:"status"` // published | draft, default: draft
}

// UpdatePostRequest 更新文章请求
type UpdatePostRequest struct {
	Title      string   `json:"title,omitempty"`
	Excerpt    string   `json:"excerpt,omitempty"`
	Content    string   `json:"content,omitempty"`
	CategoryID string   `json:"categoryId,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	ImageUrl   string   `json:"imageUrl,omitempty"`
	Status     string   `json:"status,omitempty"`
}

// PaginatedPostsResponse 分页文章响应
type PaginatedPostsResponse struct {
	Data       []PostResponse `json:"data"`
	Pagination Pagination     `json:"pagination"`
}

// Pagination 分页信息
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
