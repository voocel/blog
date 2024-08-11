package entity

import (
	"database/sql"
	"time"
)

type Advert struct {
	ID        int64        `gorm:"primarykey" json:"id"`
	Title     string       `gorm:"size:64" json:"title"`       // 标题
	Href      string       `gorm:"size:128" json:"href"`       // 跳转链接
	ImagesUrl string       `gorm:"size:128" json:"images_url"` // 广告图片地址
	IsShow    bool         `json:"is_show"`                    // 是否展示
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}

type AdvertReq struct {
	Title     string `json:"title" binding:"required" msg:"请输入标题"`
	Href      string `json:"href" binding:"required,url" msg:"跳转链接非法"`
	ImagesUrl string `json:"images_url" binding:"required,url" msg:"图片地址非法"`
	IsShow    bool   `json:"is_show"`
}
