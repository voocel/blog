package entity

import (
	"database/sql"
	"time"
)

type Tag struct {
	ID        int64        `gorm:"primarykey"`
	Name      string       `gorm:"size:32;uniqueIndex" json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}

type TagReq struct {
	ID   int64  `json:"id"`
	Name string `json:"title" binding:"required" msg:"请输入标题"`
}
