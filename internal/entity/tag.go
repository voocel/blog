package entity

import (
	"database/sql"
	"time"
)

type Tag struct {
	ID           int64        `gorm:"primarykey" json:"id"`
	Name         string       `gorm:"size:50;not null;uniqueIndex" json:"name"`
	Slug         string       `gorm:"size:100;not null;uniqueIndex" json:"slug"`
	Description  string       `gorm:"size:500" json:"description"`
	Color        string       `gorm:"size:20" json:"color"`
	ArticleCount int          `gorm:"default:0" json:"articleCount"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	DeletedAt    sql.NullTime `gorm:"index" json:"-"`

	// 关联
	Articles []Article `gorm:"many2many:article_tags" json:"articles,omitempty"`
}

// TagRequest 标签请求
type TagRequest struct {
	Name        string `json:"name" binding:"required" msg:"标签名称不能为空"`
	Slug        string `json:"slug" binding:"required" msg:"标签别名不能为空"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

// TagResponse 标签响应
type TagResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Description  string `json:"description,omitempty"`
	Color        string `json:"color,omitempty"`
	ArticleCount int    `json:"articleCount"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}
