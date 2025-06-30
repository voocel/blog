package entity

import (
	"database/sql"
	"time"
)

type Category struct {
	ID           int64        `gorm:"primarykey" json:"id"`
	Name         string       `gorm:"size:100;not null" json:"name"`
	Slug         string       `gorm:"size:150;not null;uniqueIndex" json:"slug"`
	Description  string       `gorm:"size:500" json:"description"`
	ParentID     *int64       `json:"parentId"`
	ArticleCount int          `gorm:"default:0" json:"articleCount"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	DeletedAt    sql.NullTime `gorm:"index" json:"-"`

	// 关联
	Parent   *Category  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children []Category `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Articles []Article  `gorm:"foreignKey:CategoryID" json:"articles,omitempty"`
}

// CategoryRequest 分类请求
type CategoryRequest struct {
	Name        string `json:"name" binding:"required" msg:"分类名称不能为空"`
	Slug        string `json:"slug" binding:"required" msg:"分类别名不能为空"`
	Description string `json:"description"`
	ParentID    *int64 `json:"parentId"`
}

// CategoryResponse 分类响应
type CategoryResponse struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	Slug         string             `json:"slug"`
	Description  string             `json:"description,omitempty"`
	ParentID     string             `json:"parentId,omitempty"`
	ArticleCount int                `json:"articleCount"`
	Children     []CategoryResponse `json:"children,omitempty"`
	CreatedAt    string             `json:"createdAt"`
	UpdatedAt    string             `json:"updatedAt"`
}
