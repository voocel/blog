package entity

import "time"

type Category struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}
