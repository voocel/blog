package entity

import (
	"blog/internal/entity/ctype"
	"time"
)

type Banner struct {
	ID        int64           `gorm:"primarykey" json:"id"`
	Path      string          `json:"path"`
	Hash      string          `json:"hash"`
	Name      string          `gorm:"size:38" json:"name"`
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt time.Time       `json:"-"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type BannerReq struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}
