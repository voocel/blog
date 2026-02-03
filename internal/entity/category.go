package entity

import (
	"time"
)

type Category struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	Slug      string    `gorm:"type:varchar(150);uniqueIndex;not null" json:"slug"`
	Count     int       `gorm:"type:int;default:0" json:"count"` // number of posts
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

type CategoryResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Count int    `json:"count"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug"` // optional, auto-generated if empty
}

func (Category) TableName() string {
	return "categories"
}
