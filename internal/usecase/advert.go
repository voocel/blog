package usecase

import (
	"blog/internal/entity"
	"context"
	"github.com/jinzhu/copier"
)

type AdvertUseCase struct {
	repo AdvertRepo
}

func NewAdvertUseCase(r AdvertRepo) *AdvertUseCase {
	return &AdvertUseCase{repo: r}
}

func (c *AdvertUseCase) AddAdvert(ctx context.Context, req *entity.AdvertReq) error {
	advert := new(entity.Advert)
	copier.Copy(advert, req)
	return c.repo.AddAdvertRepo(ctx, advert)
}

func (c *AdvertUseCase) Detail(ctx context.Context, id int64) (*entity.Advert, error) {
	return c.repo.AddAdvertRepo()
}

func (c *AdvertUseCase) List(ctx context.Context) ([]*entity.Advert, error) {
	return c.repo.GetAdvertListRepo(ctx)
}

func (c *AdvertUseCase) UpdateAdvert(ctx context.Context, req *entity.AdvertReq) error {
	advert := new(entity.Advert)
	copier.Copy(advert, req)
	return c.repo.UpdateAdvertRepo(ctx, advert)
}

func (c *AdvertUseCase) DeleteAdvert(ctx context.Context, id int64) error {
	return c.repo.DeleteAdvertRepo(ctx, id)
}
