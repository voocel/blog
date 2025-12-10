package entity

import "time"

// Comment represents a top-level comment or a single-level reply (no deep nesting).
type Comment struct {
	ID        string     `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PostID    string     `gorm:"type:uuid;not null;index" json:"postId"`
	UserID    string     `gorm:"type:uuid;not null;index" json:"userId"`
	ParentID  *string    `gorm:"type:uuid;index" json:"parentId,omitempty"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}

type CreateCommentRequest struct {
	Content  string  `json:"content" binding:"required"`
	ParentID *string `json:"parentId"`
}

type CommentUser struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type CommentResponse struct {
	ID          string            `json:"id"`
	ParentID    *string           `json:"parentId,omitempty"`
	Content     string            `json:"content"`
	CreatedAt   time.Time         `json:"createdAt"`
	User        CommentUser       `json:"user"`
	ReplyToUser *CommentUser      `json:"replyToUser,omitempty"`
	Replies     []CommentResponse `json:"replies,omitempty"`
}

type PaginatedCommentsResponse struct {
	Data       []CommentResponse `json:"data"`
	Pagination Pagination        `json:"pagination"`
}

func (Comment) TableName() string {
	return "comments"
}
