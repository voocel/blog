package entity

import "blog/internal/entity/ctype"

type Article struct {
	ID        int64  `json:"id" structs:"id"`
	CreatedAt string `json:"created_at" structs:"created_at"`
	UpdatedAt string `json:"updated_at" structs:"updated_at"`

	Title    string `json:"title" structs:"title"`
	Keyword  string `json:"keyword,omit(list)" structs:"keyword"`
	Abstract string `json:"abstract" structs:"abstract"`
	Content  string `json:"content,omit(list)" structs:"content"`

	LookCount     int `json:"look_count" structs:"look_count"`
	CommentCount  int `json:"comment_count" structs:"comment_count"`
	DiggCount     int `json:"digg_count" structs:"digg_count"`
	CollectsCount int `json:"collects_count" structs:"collects_count"`

	UserID       uint   `json:"user_id" structs:"user_id"`
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"`
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`

	Category string `json:"category" structs:"category"`
	Source   string `json:"source" structs:"source"`
	Link     string `json:"link" structs:"link"`

	BannerID  uint   `json:"banner_id" structs:"banner_id"`
	BannerUrl string `json:"banner_url" structs:"banner_url"`

	Tags ctype.Array `json:"tags" structs:"tags"`
}

type ArticleReq struct {
	Title    string      `json:"title" binding:"required" msg:"文章标题必填"`
	Abstract string      `json:"abstract"`
	Content  string      `json:"content" binding:"required" msg:"文章内容必填"`
	Category string      `json:"category"`
	Source   string      `json:"source"`
	Link     string      `json:"link"`
	BannerID uint        `json:"banner_id"`
	Tags     ctype.Array `json:"tags"`
}
