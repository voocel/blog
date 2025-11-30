package entity

import (
	"time"
)

// Category 分类模型
type Category struct {
	ID        string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	Slug      string    `gorm:"type:varchar(150);uniqueIndex;not null" json:"slug"`
	Count     int       `gorm:"type:int;default:0" json:"count"` // number of posts
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

// CategoryResponse 分类响应
type CategoryResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Count int    `json:"count"`
}

// CreateCategoryRequest 创建分类请求
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug"` // optional, auto-generated if empty
}

func (Category) TableName() string {
	return "categories"
}
