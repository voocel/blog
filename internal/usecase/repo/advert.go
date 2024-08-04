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
	return a.db.Create(advert).Error
}

func (a AdvertRepo) DetailRepo(ctx context.Context, id int64) (*entity.Advert, error) {
	var model = new(entity.Advert)
	err := a.db.First(model, id).Error
	return model, err
}

func (a AdvertRepo) GetAdvertListRepo(ctx context.Context) ([]*entity.Advert, error) {
	var adverts []*entity.Advert
	err := a.db.Find(&adverts).Error
	return adverts, err
}

func (a AdvertRepo) UpdateAdvertRepo(ctx context.Context, advert *entity.Advert) error {
	return a.db.Model(&entity.Advert{}).Updates(advert).Error
}

func (a AdvertRepo) DeleteAdvertRepo(ctx context.Context, id int64) error {
	return a.db.Where("id = ?", id).Delete(&entity.Advert{}).Error
}

func (a AdvertRepo) DeleteAdvertBatchRepo(ctx context.Context, ids []int64) error {
	//return a.db.Where("id in (?)", ids).Delete(&entity.Advert{}).Error
	return a.db.Delete(&entity.Advert{}, ids).Error
}
