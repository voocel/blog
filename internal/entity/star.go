package entity

import (
	"database/sql"
	"time"
)

type Star struct {
	ID        int64 `gorm:"primarykey"`
	UserID    int64 `json:"user_id"`
	ArticleID int64 `json:"article_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
