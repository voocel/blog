package entity

import (
	"database/sql"
	"time"
)

type Menu struct {
	ID           int64        `gorm:"primarykey"`
	Title        string       `gorm:"size:32" json:"title"`        // 标题
	Path         string       `gorm:"size:256" json:"path"`        // 路径
	Slogan       string       `gorm:"size:64" json:"slogan"`       // 标语
	Abstract     Array        `gorm:"type:string" json:"abstract"` // 简介
	AbstractTime int          `json:"abstract_time"`               // 简介的切换时间
	Banners      Array        `gorm:"type:string" json:"banners"`  // 菜单的图片列表
	BannerTime   int          `json:"banner_time"`                 // 菜单图片的切换时间 为 0 表示不切换
	Sort         int          `gorm:"size:10" json:"sort"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    sql.NullTime `gorm:"index" json:"-"`
}

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuReq struct {
	ID            int64       `json:"id"`
	Title         string      `json:"title"  binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path"  binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      Array       `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`
	BannerTime    int         `json:"banner_time" structs:"banner_time"`
	Sort          int         `json:"sort"  structs:"sort"`
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`
}
