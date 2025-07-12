package entity

import (
	"database/sql"
	"time"
)

type Tag struct {
	ID           int64        `gorm:"primarykey" json:"id"`
	Name         string       `gorm:"size:50;not null;uniqueIndex" json:"name"`
	Title        string       `gorm:"size:100" json:"title"`
	Description  string       `gorm:"size:500" json:"description"`
	Color        string       `gorm:"size:20" json:"color"`
	ArticleCount int          `gorm:"default:0" json:"articleCount"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	DeletedAt    sql.NullTime `gorm:"index" json:"-"`
}

// TagRequest 标签请求
type TagRequest struct {
	Name        string `json:"name" binding:"required" msg:"标签名称不能为空"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

// TagResponse 标签响应
type TagResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Color        string `json:"color,omitempty"`
	ArticleCount int    `json:"articleCount"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}
