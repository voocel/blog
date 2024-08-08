package entity

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID              int64  `gorm:"primarykey"`
	ParentCommentID int64  `json:"parent_comment_id"`                       // 父评论id
	Content         string `gorm:"size:256"`                                // 评论内容
	LikeCounts      int    `gorm:"size:8;default:0;" json:"like_counts"`    // 点赞数
	CommentCounts   int    `gorm:"size:8;default:0;" json:"comment_counts"` // 评论数
	ArticleID       string `gorm:"size:32" json:"article_id"`               // 文章id
	UserID          int64  `json:"user_id"`                                 // 用户id
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime `gorm:"index"`
}

type CommentReq struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID int64  `json:"parent_comment_id"`
}
