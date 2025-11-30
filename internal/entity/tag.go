package entity

import (
	"time"
)

// Tag 标签模型
type Tag struct {
	ID        string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

// TagResponse 标签响应
type TagResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateTagRequest 创建标签请求
type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

func (Tag) TableName() string {
	return "tags"
}
