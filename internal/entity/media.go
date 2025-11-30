package entity

import (
	"time"
)

// Media media file model
type Media struct {
	ID        string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	URL       string    `gorm:"type:varchar(500);not null" json:"url"` // Absolute URL
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // Original filename
	Type      string    `gorm:"type:varchar(20);not null" json:"type"` // image | video | document
	Size      int64     `gorm:"type:bigint" json:"size,omitempty"` // File size in bytes
	MimeType  string    `gorm:"type:varchar(100)" json:"mimeType,omitempty"`
	Path      string    `gorm:"type:varchar(500)" json:"path,omitempty"` // Server file path
	Date      string    `gorm:"type:varchar(30);not null" json:"date"` // ISO Date
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

// MediaResponse media file response
type MediaResponse struct {
	ID   string `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
	Type string `json:"type"`
	Date string `json:"date"`
}

func (Media) TableName() string {
	return "media"
}
