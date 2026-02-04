package entity

import "time"

// Comment represents a top-level comment or a single-level reply (no deep nesting).
type Comment struct {
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    int64      `gorm:"not null;index" json:"postId"`
	UserID    int64      `gorm:"not null;index" json:"userId"`
	ParentID  *int64     `gorm:"index" json:"parentId,omitempty"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}

type CreateCommentRequest struct {
	Content  string `json:"content" binding:"required"`
	ParentID *int64 `json:"parentId"`
}

type CommentUser struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type CommentResponse struct {
	ID          int64             `json:"id"`
	ParentID    *int64            `json:"parentId,omitempty"`
	Content     string            `json:"content"`
	CreatedAt   time.Time         `json:"createdAt"`
	User        CommentUser       `json:"user"`
	ReplyToUser *CommentUser      `json:"replyToUser,omitempty"`
	Replies     []CommentResponse `json:"replies,omitempty"`
}

// AdminCommentResponse is used by admin moderation endpoints.
type AdminCommentResponse struct {
	ID        int64       `json:"id"`
	ParentID  *int64      `json:"parentId"`
	Content   string      `json:"content"`
	CreatedAt time.Time   `json:"createdAt"`
	PostID    int64       `json:"postId"`
	PostTitle string      `json:"postTitle"`
	User      CommentUser `json:"user"`
}

type PaginatedCommentsResponse struct {
	Data       []CommentResponse `json:"data"`
	Pagination Pagination        `json:"pagination"`
}

func (Comment) TableName() string {
	return "comments"
}
