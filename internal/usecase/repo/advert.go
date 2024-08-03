package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

type AdvertRepo struct {
	db *gorm.DB
}

func NewAdvertRepo(db *gorm.DB) *AdvertRepo {
	return &AdvertRepo{db: db}
}

func (a AdvertRepo) AddAdvertRepo(ctx context.Context, advert *entity.Advert) error {
	//TODO implement me
	panic("implement me")
}

func (a AdvertRepo) DetailRepo(ctx context.Context, id int64) (*entity.Advert, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdvertRepo) GetAdvertListRepo(ctx context.Context) ([]*entity.Advert, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdvertRepo) UpdateAdvertRepo(ctx context.Context, advert *entity.Advert) error {
	//TODO implement me
	panic("implement me")
}

func (a AdvertRepo) DeleteAdvertRepo(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
