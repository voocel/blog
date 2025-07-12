package entity

import (
	"database/sql"
	"time"
)

// Discussion 讨论实体
type Discussion struct {
	ID         int64        `gorm:"primarykey" json:"id"`
	Title      string       `gorm:"size:255;not null" json:"title"`
	Content    string       `gorm:"type:text" json:"content"`
	Status     string       `gorm:"size:20;default:active" json:"status"` // active, inactive
	ViewCount  int          `gorm:"default:0" json:"viewCount"`
	ReplyCount int          `gorm:"default:0" json:"replyCount"`
	UserID     int64        `json:"userId"`
	Tags       Array        `json:"tags"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
	DeletedAt  sql.NullTime `gorm:"index" json:"-"`
}

// Reply 回复实体
type Reply struct {
	ID           int64        `gorm:"primarykey" json:"id"`
	Content      string       `gorm:"type:text" json:"content"`
	DiscussionID int64        `json:"discussionId"`
	UserID       int64        `json:"userId"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	DeletedAt    sql.NullTime `gorm:"index" json:"-"`
}

// DiscussionRequest 创建讨论请求
type DiscussionRequest struct {
	Title   string  `json:"title" binding:"required" msg:"标题不能为空"`
	Content string  `json:"content" binding:"required" msg:"内容不能为空"`
	TagIds  []int64 `json:"tagIds"`
	Status  string  `json:"status"`
}

// DiscussionResponse 讨论响应
type DiscussionResponse struct {
	ID         int64           `json:"id"`
	Title      string          `json:"title"`
	Content    string          `json:"content"`
	Status     string          `json:"status"`
	ViewCount  int             `json:"viewCount"`
	ReplyCount int             `json:"replyCount"`
	Tags       []TagResponse   `json:"tags"`
	Author     AuthorResponse  `json:"author"`
	Replies    []ReplyResponse `json:"replies,omitempty"`
	CreatedAt  string          `json:"createdAt"`
	UpdatedAt  string          `json:"updatedAt"`
}

// ReplyResponse 回复响应
type ReplyResponse struct {
	ID        int64          `json:"id"`
	Content   string         `json:"content"`
	Author    AuthorResponse `json:"author"`
	CreatedAt string         `json:"createdAt"`
}

// AuthorResponse 作者响应
type AuthorResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// DiscussionDetailResponse 讨论详情响应
type DiscussionDetailResponse struct {
	ID         int64           `json:"id"`
	Title      string          `json:"title"`
	Content    string          `json:"content"`
	Status     string          `json:"status"`
	ViewCount  int             `json:"viewCount"`
	ReplyCount int             `json:"replyCount"`
	Tags       []TagResponse   `json:"tags"`
	Author     AuthorResponse  `json:"author"`
	Replies    []ReplyResponse `json:"replies"`
	CreatedAt  string          `json:"createdAt"`
	UpdatedAt  string          `json:"updatedAt"`
}
