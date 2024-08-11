package entity

import (
	"database/sql"
	"time"
)

type Star struct {
	ID        int64        `gorm:"primarykey" json:"id"`
	UserID    int64        `json:"user_id"`
	ArticleID int64        `json:"article_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}
