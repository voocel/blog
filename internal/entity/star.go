package entity

import "time"

type Star struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	UserID    int64  `json:"user_id"`
	ArticleID string `gorm:"size:64"`
	CreatedAt time.Time
}
