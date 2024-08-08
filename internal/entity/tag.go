package entity

import (
	"database/sql"
	"time"
)

type Tag struct {
	ID        int64  `gorm:"primarykey"`
	Name      string `gorm:"size:32;uniqueIndex" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
