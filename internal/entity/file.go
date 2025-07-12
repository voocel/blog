package entity

import (
	"database/sql"
	"time"
)

// File 文件实体
type File struct {
	ID           int64        `gorm:"primarykey" json:"id"`
	Filename     string       `gorm:"size:255;not null" json:"filename"`
	OriginalName string       `gorm:"size:255;not null" json:"originalName"`
	MimeType     string       `gorm:"size:100" json:"mimeType"`
	Size         int64        `json:"size"`
	URL          string       `gorm:"size:500" json:"url"`
	Path         string       `gorm:"size:500" json:"path"`
	Type         string       `gorm:"size:20;default:file" json:"type"` // file, folder
	UserID       int64        `json:"userId"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	DeletedAt    sql.NullTime `gorm:"index" json:"-"`
}

// FileResponse 文件响应
type FileResponse struct {
	ID           int64  `json:"id"`
	Filename     string `json:"filename"`
	OriginalName string `json:"originalName"`
	MimeType     string `json:"mimeType"`
	Size         int64  `json:"size"`
	URL          string `json:"url"`
	Path         string `json:"path"`
	Type         string `json:"type"`
	CreatedAt    string `json:"createdAt"`
}

// CreateFolderRequest 创建文件夹请求
type CreateFolderRequest struct {
	Name string `json:"name" binding:"required" msg:"文件夹名称不能为空"`
	Path string `json:"path" binding:"required" msg:"路径不能为空"`
}
