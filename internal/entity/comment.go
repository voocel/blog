package entity

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID           int64        `gorm:"primarykey" json:"id"`
	Content      string       `gorm:"type:text;not null" json:"content"`
	Status       string       `gorm:"size:20;default:pending" json:"status"` // approved, pending, rejected
	ArticleID    *int64       `json:"articleId"`
	DiscussionID *int64       `json:"discussionId"`
	ParentID     *int64       `json:"parentId"` // 父评论ID（回复）
	UserID       int64        `gorm:"not null" json:"userId"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	DeletedAt    sql.NullTime `gorm:"index" json:"-"`
}

// CommentRequest 评论请求
type CommentRequest struct {
	Content      string `json:"content" binding:"required" msg:"评论内容不能为空"`
	ArticleID    *int64 `json:"articleId"`
	DiscussionID *int64 `json:"discussionId"`
	ParentID     *int64 `json:"parentId"`
}

// CommentResponse 评论响应
type CommentResponse struct {
	ID           int64             `json:"id"`
	Content      string            `json:"content"`
	Status       string            `json:"status"`
	ArticleID    int64             `json:"articleId,omitempty"`
	DiscussionID int64             `json:"discussionId,omitempty"`
	ParentID     int64             `json:"parentId,omitempty"`
	Author       AuthorResponse    `json:"author"`
	Replies      []CommentResponse `json:"replies,omitempty"`
	CreatedAt    string            `json:"createdAt"`
	UpdatedAt    string            `json:"updatedAt"`
}
