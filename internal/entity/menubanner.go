package entity

import (
	"database/sql"
	"time"
)

type MenuBanner struct {
	ID        int64        `gorm:"primarykey" json:"id"`
	MenuID    int64        `json:"menu_id"`
	BannerID  int64        `json:"banner_id"`
	Sort      int          `gorm:"size:10" json:"sort"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}
