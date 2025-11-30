package usecase

import (
	"blog/internal/entity"
	"blog/pkg/jwt"
	"blog/pkg/util"
	"context"
	"errors"
	"strings"
)

type AuthUseCase struct {
	userRepo UserRepo
}

func NewAuthUseCase(userRepo UserRepo) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

// Login user login
func (uc *AuthUseCase) Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error) {
	user, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		Token: token,
		User: entity.UserResponse{
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
			Avatar:   user.Avatar,
		},
	}, nil
}

// GetCurrentUser get current user
func (uc *AuthUseCase) GetCurrentUser(ctx context.Context, userID string) (*entity.UserResponse, error) {
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &entity.UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Avatar:   user.Avatar,
	}, nil
}

// Register user registration
func (uc *AuthUseCase) Register(ctx context.Context, req entity.RegisterRequest) (*entity.LoginResponse, error) {
	existUser, _ := uc.userRepo.GetByEmail(ctx, req.Email)
	if existUser != nil {
		return nil, errors.New("email already exists")
	}

	username := strings.TrimSpace(req.Username)
	if username == "" {
		username = util.GenerateRandomUsername()
	}

	for {
		existByUsername, _ := uc.userRepo.GetByUsername(ctx, username)
		if existByUsername == nil {
			break // Username available
		}
		username = util.GenerateRandomUsername()
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "visitor", // Default role is visitor
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		Token: token,
		User: entity.UserResponse{
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
			Avatar:   user.Avatar,
		},
	}, nil
}
