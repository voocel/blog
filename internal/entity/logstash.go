package entity

import (
	"encoding/json"
	"time"
)

type Logstash struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	IP        string    `gorm:"size:32" json:"ip"`
	Addr      string    `gorm:"size:64" json:"addr"`
	Level     Level     `gorm:"size:4" json:"level"`
	Content   string    `gorm:"size:128" json:"content"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func (s Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())

}

func (s Level) String() string {
	switch s {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	default:
		return "other"
	}
}
