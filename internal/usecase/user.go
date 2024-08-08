package usecase

import (
	"blog/internal/entity"
	"blog/pkg/util"
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

type UserUseCase struct {
	repo UserRepo
	sf   singleflight.Group
}

func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (u *UserUseCase) UserLogin(ctx context.Context, req entity.UserLoginReq) (*entity.User, error) {
	user, err := u.repo.GetUserByNameRepo(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user.Status != 1 {
		return nil, errors.New("this user is disable")
	}
	if !util.VerifyPassword(req.Password, user.Password) {
		return nil, errors.New("password incorrect")
	}
	return user, nil
}

func (u *UserUseCase) UserRegister(ctx context.Context, req entity.UserRegisterReq) error {
	exist, err := u.repo.GetUserByNameExistRepo(ctx, req.Username)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("username already exists")
	}
	userInfo := &entity.User{}
	userInfo.Password, err = util.EncryptPassword(req.Password)
	if err != nil {
		return err
	}
	userInfo.Username = req.Username
	userInfo.Nickname = "unknown"
	userInfo.Role = 1
	userInfo.Status = 1
	userInfo.Birthday = time.Now()
	userInfo.LastLoginTime = time.Now()
	return u.repo.AddUserRepo(ctx, userInfo)
}

func (u *UserUseCase) GetUserById(ctx context.Context, uid int64) (*entity.User, error) {
	v, err, _ := u.sf.Do("key1", func() (interface{}, error) {
		userInfo, err := u.repo.GetUserByIdRepo(ctx, uid)
		if err != nil {
			return nil, fmt.Errorf("GetUserByIdRepo err: %w", err)
		}
		// todo set cache
		return userInfo, err
	})

	return v.(*entity.User), err
}
