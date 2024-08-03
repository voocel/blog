package entity

import "time"

type Advert struct {
	ID        int64     `gorm:"primarykey" json:"id,select($any)"  structs:"-"`
	Title     string    `gorm:"size:32" json:"title"`
	Href      string    `json:"href"`
	Images    string    `json:"images"`
	IsShow    bool      `json:"is_show"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

type AdvertReq struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`
	Images string `json:"images" binding:"required,url" msg:"图片地址非法" structs:"images"`
	IsShow bool   `json:"is_show" structs:"is_show"`
}
