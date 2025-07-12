package entity

import (
	"database/sql"
	"time"
)

type Category struct {
	ID           int64        `gorm:"primarykey" json:"id"`
	Name         string       `gorm:"size:100;not null" json:"name"`
	Path         string       `gorm:"size:150;not null;uniqueIndex" json:"path"`
	Description  string       `gorm:"size:500" json:"description"`
	ArticleCount int          `gorm:"default:0" json:"articleCount"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	DeletedAt    sql.NullTime `gorm:"index" json:"-"`
}

// CategoryRequest 分类请求
type CategoryRequest struct {
	Name        string `json:"name" binding:"required" msg:"分类名称不能为空"`
	Path        string `json:"path" binding:"required" msg:"分类路径不能为空"`
	Description string `json:"description"`
}

// CategoryResponse 分类响应
type CategoryResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Description  string `json:"description,omitempty"`
	ArticleCount int    `json:"articleCount"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}
