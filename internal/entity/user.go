package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID            int64        `gorm:"primarykey" json:"id"`
	Username      string       `gorm:"size:32" json:"username"`
	Password      string       `gorm:"size:64" json:"password"`
	Mobile        string       `gorm:"size:32" json:"mobile"`
	Nickname      string       `gorm:"size:32" json:"nickname"`
	Email         string       `gorm:"size:32" json:"email"`
	Avatar        string       `gorm:"size:128" json:"avatar"`
	Summary       string       `gorm:"size:256" json:"summary,omitempty"`
	IP            string       `gorm:"size:20" json:"ip"`
	Scores        int64        `json:"scores"` // 积分
	Sex           int8         `json:"sex"`
	Role          int8         `json:"role"`   // 1 管理员 0 普通用户
	Source        int8         `json:"source"` // 注册来源
	Status        int8         `json:"status"`
	Birthday      time.Time    `json:"-"`
	LastLoginTime time.Time    `json:"last_login_time"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     sql.NullTime `gorm:"index" json:"-"`
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
