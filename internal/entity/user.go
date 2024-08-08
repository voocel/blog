package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID            int64     `gorm:"primarykey"`
	Username      string    `gorm:"size:32" json:"username"`
	Password      string    `gorm:"size:64" json:"password"`
	Mobile        string    `gorm:"size:32" json:"mobile"`
	Nickname      string    `gorm:"size:32" json:"nickname"`
	Email         string    `gorm:"size:32" json:"email"`
	Avatar        string    `gorm:"size:128" json:"avatar"`
	Summary       string    `gorm:"size:256" json:"summary,omitempty"`
	Sex           int8      `json:"sex"`
	Role          int8      `json:"role"`
	Status        int8      `json:"status"`
	Birthday      time.Time `json:"-"`
	LastLoginTime time.Time `json:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//Mobile   string `json:"mobile,omitempty"`
	//Nickname string `json:"nickname,omitempty"`
}
