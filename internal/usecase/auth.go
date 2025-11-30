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

// Login 用户登录
func (uc *AuthUseCase) Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error) {
	user, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// 验证密码
	if !util.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// 生成 JWT token
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

// GetCurrentUser 获取当前用户
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

// Register 用户注册
func (uc *AuthUseCase) Register(ctx context.Context, req entity.RegisterRequest) (*entity.LoginResponse, error) {
	// 检查邮箱是否已存在
	existUser, _ := uc.userRepo.GetByEmail(ctx, req.Email)
	if existUser != nil {
		return nil, errors.New("email already exists")
	}

	// 生成用户名（如果未提供）
	username := strings.TrimSpace(req.Username)
	if username == "" {
		username = util.GenerateRandomUsername()
	}

	// 检查用户名是否已存在（循环直到找到唯一用户名）
	for {
		existByUsername, _ := uc.userRepo.GetByUsername(ctx, username)
		if existByUsername == nil {
			break // 用户名可用
		}
		// 如果用户名已存在，重新生成
		username = util.GenerateRandomUsername()
	}

	// 加密密码
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 创建新用户
	user := &entity.User{
		Username: username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "visitor", // 默认角色为 visitor
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	// 生成 JWT token
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
