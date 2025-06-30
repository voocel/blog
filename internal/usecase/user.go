package usecase

import (
	"blog/internal/entity"
	"context"

	"golang.org/x/sync/singleflight"
)

type UserUseCase struct {
	repo UserRepo
	sf   singleflight.Group
}

func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (u *UserUseCase) GetUserById(ctx context.Context, uid int64) (*entity.User, error) {
	return u.repo.GetUserByIdRepo(ctx, uid)
}

func (u *UserUseCase) UserList(ctx context.Context, page, pageSize int, search string) ([]*entity.User, int64, error) {
	return u.repo.GetUsersRepo(ctx, page, pageSize, search)
}
