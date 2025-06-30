package entity

import (
	"database/sql"
	"time"
)

type Article struct {
	ID           int64        `gorm:"primarykey" json:"id"`
	Title        string       `gorm:"size:255;not null" json:"title"`
	Subtitle     string       `gorm:"size:255" json:"subtitle"`
	Content      string       `gorm:"type:text;not null" json:"content"`
	Excerpt      string       `gorm:"size:500" json:"excerpt"`
	CoverImage   string       `gorm:"size:500" json:"coverImage"`
	Status       string       `gorm:"size:20;default:draft" json:"status"` // published, draft
	IsOriginal   bool         `gorm:"default:true" json:"isOriginal"`
	ViewCount    int          `gorm:"default:0" json:"viewCount"`
	LikeCount    int          `gorm:"default:0" json:"likeCount"`
	CommentCount int          `gorm:"default:0" json:"commentCount"`
	UserID       int64        `gorm:"not null" json:"userId"`
	CategoryID   int64        `json:"categoryId"`
	Source       string       `gorm:"size:100" json:"source"`
	Link         string       `gorm:"size:500" json:"link"`
	PublishedAt  *time.Time   `json:"publishedAt"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	DeletedAt    sql.NullTime `gorm:"index" json:"-"`

	// 关联
	User     *User     `gorm:"foreignKey:UserID" json:"author,omitempty"`
	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Tags     []Tag     `gorm:"many2many:article_tags" json:"tags,omitempty"`
}

// ArticleRequest 创建文章请求
type ArticleRequest struct {
	Title       string  `json:"title" binding:"required" msg:"文章标题必填"`
	Subtitle    string  `json:"subtitle"`
	Content     string  `json:"content" binding:"required" msg:"文章内容必填"`
	Excerpt     string  `json:"excerpt"`
	CoverImage  string  `json:"coverImage"`
	CategoryID  int64   `json:"categoryId" binding:"required" msg:"分类不能为空"`
	TagIds      []int64 `json:"tagIds"`
	Status      string  `json:"status"`
	IsOriginal  bool    `json:"isOriginal"`
	Source      string  `json:"source"`
	Link        string  `json:"link"`
	PublishedAt string  `json:"publishedAt"`
}

// ArticleUpdateRequest 更新文章请求
type ArticleUpdateRequest struct {
	Title       string  `json:"title"`
	Subtitle    string  `json:"subtitle"`
	Content     string  `json:"content"`
	Excerpt     string  `json:"excerpt"`
	CoverImage  string  `json:"coverImage"`
	CategoryID  int64   `json:"categoryId"`
	TagIds      []int64 `json:"tagIds"`
	Status      string  `json:"status"`
	IsOriginal  bool    `json:"isOriginal"`
	Source      string  `json:"source"`
	Link        string  `json:"link"`
	PublishedAt string  `json:"publishedAt"`
}

// ArticleResponse 文章响应
type ArticleResponse struct {
	ID           string           `json:"id"`
	Title        string           `json:"title"`
	Subtitle     string           `json:"subtitle,omitempty"`
	Content      string           `json:"content"`
	Excerpt      string           `json:"excerpt"`
	CoverImage   string           `json:"coverImage,omitempty"`
	Status       string           `json:"status"`
	IsOriginal   bool             `json:"isOriginal"`
	ViewCount    int              `json:"viewCount"`
	LikeCount    int              `json:"likeCount"`
	CommentCount int              `json:"commentCount"`
	Tags         []TagResponse    `json:"tags"`
	Category     CategoryResponse `json:"category"`
	Author       AuthorResponse   `json:"author"`
	PublishedAt  string           `json:"publishedAt,omitempty"`
	CreatedAt    string           `json:"createdAt"`
	UpdatedAt    string           `json:"updatedAt"`
}
