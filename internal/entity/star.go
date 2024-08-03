package entity

import "time"

type Star struct {
	ID        int64 `gorm:"primarykey" json:"id"`
	UserID    int64 `json:"user_id"`
	ArticleID int64
	CreatedAt time.Time
}
