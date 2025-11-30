package entity

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID         string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Username   string    `gorm:"type:varchar(50);not null" json:"username"` // 昵称，非唯一
	Email      string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"` // 唯一标识
	Password   string    `gorm:"type:varchar(255)" json:"-"` // 可选，OAuth 用户可为空
	Role       string    `gorm:"type:varchar(20);not null;default:'visitor'" json:"role"` // admin | visitor
	Avatar     string    `gorm:"type:varchar(500)" json:"avatar,omitempty"`
	Bio        string    `gorm:"type:text" json:"bio,omitempty"`
	Provider   string    `gorm:"type:varchar(20);not null;default:'email';uniqueIndex:idx_provider_user" json:"provider"` // email | google | github | apple
	ProviderID string    `gorm:"type:varchar(255);uniqueIndex:idx_provider_user" json:"-"` // 第三方平台用户ID，与provider组合唯一
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// BeforeCreate GORM 钩子，在创建前确保 provider 有值
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Provider == "" {
		u.Provider = "email"
	}
	return nil
}

// UserResponse 用户响应（不包含密码）
type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar,omitempty"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username"` // 可选，如果为空则自动生成
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// UpdateProfileRequest 更新个人信息请求
type UpdateProfileRequest struct {
	Username string `json:"username,omitempty"`
	Bio      string `json:"bio,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}
