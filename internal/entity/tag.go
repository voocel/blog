package entity

import (
	"time"
)

type Tag struct {
	ID        string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

type TagResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

func (Tag) TableName() string {
	return "tags"
}
