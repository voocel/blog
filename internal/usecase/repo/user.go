package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u UserRepo) GetUserByIdRepo(ctx context.Context, uid int64) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) GetUserByNameRepo(ctx context.Context, name string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) GetUserByNameExistRepo(ctx context.Context, name string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) GetUsersRepo(ctx context.Context) ([]*entity.User, int, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) AddUserRepo(ctx context.Context, user *entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) UpdateAvatarUserRepo(ctx context.Context, uid int64, avatar string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) UpdateAddressUserRepo(ctx context.Context, uid int64, address string) (int, error) {
	//TODO implement me
	panic("implement me")
}
