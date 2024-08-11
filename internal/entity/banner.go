package entity

import (
	"database/sql"
	"time"
)

type Banner struct {
	ID            int64        `gorm:"primarykey" json:"id"`
	Path          string       `gorm:"size:256;not null;default:''" json:"path"` // 存储路径
	Hash          string       `gorm:"size:256;not null;default:''" json:"hash"` // 图片hash,用于判重
	Name          string       `gorm:"size:64;not null;default:''" json:"name"`  // 图片名称
	StorageMethod int8         `gorm:"default:0" json:"storage_method"`          // 存储类型 0 本地
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     sql.NullTime `gorm:"index" json:"-"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type BannerUpdateReq struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}
