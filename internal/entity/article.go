package entity

import (
	"database/sql"
	"time"
)

type Article struct {
	ID            int64  `gorm:"primarykey" json:"id"`
	Title         string `gorm:"size:32" json:"title"`                     // 标题
	Keyword       string `gorm:"size:32" json:"keyword"`                   // 关键字
	Abstract      string `gorm:"size:32" json:"abstract"`                  // 文章摘要
	Content       string `gorm:"size:32" json:"content"`                   // 文章内容
	ViewCounts    int    `gorm:"size:8;default:0;" json:"view_counts"`     // 浏览量
	CommentCounts int    `gorm:"size:8;default:0;" json:"comment_counts"`  // 评论数
	LikeCounts    int    `gorm:"size:8;default:0;" json:"like_counts"`     // 点赞数
	StarCounts    int    `gorm:"size:8;default:0;" json:"collects_counts"` // 收藏量
	UserID        uint   `json:"user_id"`
	UserNickname  string `gorm:"size:32" json:"user_nickname"`
	UserAvatar    string `gorm:"size:256" json:"user_avatar"`
	Category      string `gorm:"size:32" json:"category"` // 分类
	Source        string `gorm:"size:32" json:"source"`   // 来源
	Link          string `gorm:"size:256" json:"link"`    // 原文链接
	BannerID      int64  `json:"banner_id"`               // 封面id
	BannerUrl     string `json:"banner_url"`              // 封面链接
	Tags          Array  `json:"tags"`                    // 标签

	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}

type ArticleReq struct {
	Title     string `json:"title" binding:"required" msg:"文章标题必填"`
	Abstract  string `json:"abstract"`
	Content   string `json:"content" binding:"required" msg:"文章内容必填"`
	Category  string `json:"category"`
	Source    string `json:"source"`
	Link      string `json:"link"`
	BannerID  int64  `json:"banner_id"`
	BannerUrl string `json:"banner_url"`
	Tags      Array  `json:"tags"`
}

type ArticleUpdateReq struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Abstract  string `json:"abstract"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	Source    string `json:"source"`
	Link      string `json:"link"`
	BannerID  int64  `json:"banner_id"`
	BannerUrl string `json:"banner_url"`
	Tags      Array  `json:"tags"`
}
