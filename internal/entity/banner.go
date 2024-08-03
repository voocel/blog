package entity

import (
	"blog/internal/entity/ctype"
	"time"
)

type Banner struct {
	ID        int64           `gorm:"primarykey" json:"id"`
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片的类型，本地还是七牛
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt time.Time       `json:"-"`
}
