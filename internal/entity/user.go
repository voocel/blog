package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID            int64        `gorm:"primarykey" json:"id"`
	Username      string       `gorm:"size:32;not null;index" json:"username"`
	Email         string       `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Password      string       `gorm:"size:255;not null" json:"password"`
	Nickname      string       `gorm:"size:50" json:"nickname"`
	Avatar        string       `gorm:"size:255" json:"avatar"`
	Website       string       `gorm:"size:255" json:"website"`
	Description   string       `gorm:"size:500" json:"description"`
	Role          string       `gorm:"size:20;default:user" json:"role"`     // admin, user
	Status        string       `gorm:"size:20;default:active" json:"status"` // active, inactive
	Mobile        string       `gorm:"size:20" json:"mobile"`
	IP            string       `gorm:"size:45" json:"ip"`
	Scores        int64        `gorm:"default:0" json:"scores"` // 积分
	Sex           int8         `gorm:"default:0" json:"sex"`
	Source        int8         `gorm:"default:0" json:"source"` // 注册来源
	Birthday      time.Time    `json:"birthday"`
	LastLoginTime time.Time    `json:"lastLoginTime"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt"`
	DeletedAt     sql.NullTime `gorm:"index" json:"-"`
}

// UserAdminResponse 管理员用户响应
type UserAdminResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"`
	Status      string `json:"status"`
	Nickname    string `json:"nickname,omitempty"`
	Website     string `json:"website,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username    string `json:"username" binding:"required,min=3,max=32" msg:"用户名长度3-32位"`
	Email       string `json:"email" binding:"required,email" msg:"邮箱格式不正确"`
	Password    string `json:"password" binding:"required,min=6" msg:"密码至少6位"`
	Role        string `json:"role,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	Website     string `json:"website,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Username    string `json:"username,omitempty"`
	Email       string `json:"email,omitempty"`
	Role        string `json:"role,omitempty"`
	Status      string `json:"status,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	Website     string `json:"website,omitempty"`
	Description string `json:"description,omitempty"`
}
