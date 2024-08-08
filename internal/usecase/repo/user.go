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
	var user = new(entity.User)
	err := u.db.First(user, uid).Error
	return user, err
}

func (u UserRepo) GetUserByNameRepo(ctx context.Context, name string) (*entity.User, error) {
	var user = new(entity.User)
	err := u.db.First(user, "username = ?", name).Error
	return user, err
}

func (u UserRepo) GetUserByNameExistRepo(ctx context.Context, name string) (bool, error) {
	u.db.Where("username = ?", name).First(&entity.User{})
	return u.db.RowsAffected > 0, u.db.Error
}

func (u UserRepo) GetUsersRepo(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	err := u.db.Find(users).Error
	return users, err
}

func (u UserRepo) AddUserRepo(ctx context.Context, user *entity.User) error {
	return u.db.Create(user).Error
}

func (u UserRepo) UpdateAvatarUserRepo(ctx context.Context, uid int64, avatar string) error {
	return u.db.Model(&entity.User{}).Where("id = ?", uid).Update("avatar", avatar).Error
}

func (u UserRepo) UpdateAddressUserRepo(ctx context.Context, uid int64, address string) error {
	return u.db.Model(&entity.User{}).Where("id = ?", uid).Update("address", address).Error
}
