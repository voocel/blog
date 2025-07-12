package entity

import (
	"database/sql"
	"time"
)

// FriendLink 友情链接实体
type FriendLink struct {
	ID          int64        `gorm:"primarykey" json:"id"`
	Name        string       `gorm:"size:100;not null" json:"name"`
	URL         string       `gorm:"size:255;not null" json:"url"`
	Logo        string       `gorm:"size:255" json:"logo"`
	Description string       `gorm:"size:500" json:"description"`
	Status      string       `gorm:"size:20;default:active" json:"status"` // active, inactive
	SortOrder   int          `gorm:"default:0" json:"sortOrder"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   sql.NullTime `gorm:"index" json:"-"`
}

// FriendLinkRequest 友链请求
type FriendLinkRequest struct {
	Name        string `json:"name" binding:"required" msg:"友链名称不能为空"`
	URL         string `json:"url" binding:"required,url" msg:"友链地址格式不正确"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Status      string `json:"status"`
	SortOrder   int    `json:"sortOrder"`
}

// FriendLinkResponse 友链响应
type FriendLinkResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Status      string `json:"status"`
	SortOrder   int    `json:"sortOrder"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
