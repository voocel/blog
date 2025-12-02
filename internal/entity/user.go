package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Username   string    `gorm:"type:varchar(50);not null" json:"username"`               // Nickname, non-unique
	Email      string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`     // Unique identifier
	Password   string    `gorm:"type:varchar(255)" json:"-"`                              // Optional, can be empty for OAuth users
	Role       string    `gorm:"type:varchar(20);not null;default:'visitor'" json:"role"` // admin | visitor
	Avatar     string    `gorm:"type:varchar(500)" json:"avatar,omitempty"`
	Bio        string    `gorm:"type:text" json:"bio,omitempty"`
	Provider   string    `gorm:"type:varchar(20);not null;default:'email';uniqueIndex:idx_provider_user" json:"provider"` // email | google | github | apple
	ProviderID string    `gorm:"type:varchar(255);uniqueIndex:idx_provider_user" json:"-"`                                // Third-party platform user ID, unique with provider
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}

// BeforeCreate GORM hook, ensure provider has a value before creation
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Provider == "" {
		u.Provider = "email"
	}
	return nil
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username"` // Optional, auto-generated if empty
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresIn    int64        `json:"expires_in"` // Access token expiration in seconds
	User         UserResponse `json:"user"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type UpdateProfileRequest struct {
	Username string `json:"username,omitempty"`
	Bio      string `json:"bio,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}
