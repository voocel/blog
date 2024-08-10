package entity

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        int64        `gorm:"primarykey"`
	Name      string       `gorm:"size:64" json:"name"`
	Status    int8         `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}
