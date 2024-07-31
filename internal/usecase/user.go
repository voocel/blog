package usecase

import "golang.org/x/sync/singleflight"

type UserUseCase struct {
	repo UserRepo
	sf   singleflight.Group
}

func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{repo: r}
}
