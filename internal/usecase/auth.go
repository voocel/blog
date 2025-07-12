package usecase

import (
	"blog/internal/entity"
	"blog/pkg/jwt"
	"blog/pkg/util"
	"context"
	"errors"
	"time"
)

// AuthUseCase 认证业务逻辑
type AuthUseCase struct {
	userRepo UserRepo
}

// NewAuthUseCase 创建认证业务逻辑实例
func NewAuthUseCase(userRepo UserRepo) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
	}
}

// Login 用户登录
func (a *AuthUseCase) Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error) {
	// 根据邮箱查找用户
	user, err := a.userRepo.GetUserByEmailRepo(ctx, req.Email)
	if err != nil {
		return nil, errors.New("邮箱或密码错误")
	}

	// 验证密码
	if !util.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("邮箱或密码错误")
	}

	// 检查用户状态
	if user.Status != "active" {
		return nil, errors.New("账户已被禁用")
	}

	// 生成JWT令牌
	token, expiresIn, err := jwt.GenerateToken(user)
	if err != nil {
		return nil, errors.New("生成令牌失败")
	}

	// 生成刷新令牌
	refreshToken, err := jwt.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, errors.New("生成刷新令牌失败")
	}

	// 更新最后登录时间
	user.LastLoginTime = time.Now()
	a.userRepo.UpdateUserRepo(ctx, user)

	// 构造响应
	userInfo := &entity.UserInfo{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Role:        user.Role,
		Nickname:    user.Nickname,
		Website:     user.Website,
		Description: user.Description,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   user.UpdatedAt.Format(time.RFC3339),
	}

	return &entity.LoginResponse{
		User:         userInfo,
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}

// Register 用户注册
func (a *AuthUseCase) Register(ctx context.Context, req entity.RegisterRequest) (*entity.LoginResponse, error) {
	// 验证密码确认
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("密码确认不匹配")
	}

	// 检查用户名是否存在
	exists, err := a.userRepo.GetUserByNameExistRepo(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否存在
	exists, err = a.userRepo.GetUserByEmailExistRepo(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("邮箱已存在")
	}

	// 加密密码
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	// 创建用户
	user := &entity.User{
		Username:      req.Username,
		Email:         req.Email,
		Password:      hashedPassword,
		Role:          "user",
		Status:        "active",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		LastLoginTime: time.Now(),
	}

	err = a.userRepo.AddUserRepo(ctx, user)
	if err != nil {
		return nil, errors.New("创建用户失败")
	}

	// 登录逻辑
	return a.Login(ctx, entity.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
}

// RefreshToken 刷新令牌
func (a *AuthUseCase) RefreshToken(ctx context.Context, req entity.RefreshTokenRequest) (*entity.RefreshTokenResponse, error) {
	// 验证刷新令牌
	claims, err := jwt.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, errors.New("无效的刷新令牌")
	}

	// 获取用户信息
	user, err := a.userRepo.GetUserByIdRepo(ctx, claims.UserID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 检查用户状态
	if user.Status != "active" {
		return nil, errors.New("账户已被禁用")
	}

	// 生成新的访问令牌
	token, expiresIn, err := jwt.GenerateToken(user)
	if err != nil {
		return nil, errors.New("生成令牌失败")
	}

	return &entity.RefreshTokenResponse{
		Token:     token,
		ExpiresIn: expiresIn,
	}, nil
}

// GetProfile 获取用户资料
func (a *AuthUseCase) GetProfile(ctx context.Context, userID int64) (*entity.UserInfo, error) {
	user, err := a.userRepo.GetUserByIdRepo(ctx, userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	return &entity.UserInfo{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Role:        user.Role,
		Nickname:    user.Nickname,
		Website:     user.Website,
		Description: user.Description,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// UpdateProfile 更新用户资料
func (a *AuthUseCase) UpdateProfile(ctx context.Context, userID int64, req entity.UpdateProfileRequest) (*entity.UserInfo, error) {
	user, err := a.userRepo.GetUserByIdRepo(ctx, userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 更新字段
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Website != "" {
		user.Website = req.Website
	}
	if req.Description != "" {
		user.Description = req.Description
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	user.UpdatedAt = time.Now()

	err = a.userRepo.UpdateUserRepo(ctx, user)
	if err != nil {
		return nil, errors.New("更新用户资料失败")
	}

	return a.GetProfile(ctx, userID)
}

// ChangePassword 修改密码
func (a *AuthUseCase) ChangePassword(ctx context.Context, userID int64, req entity.ChangePasswordRequest) error {
	user, err := a.userRepo.GetUserByIdRepo(ctx, userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 验证旧密码
	if !util.CheckPasswordHash(req.OldPassword, user.Password) {
		return errors.New("原密码错误")
	}

	// 加密新密码
	hashedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("密码加密失败")
	}

	user.Password = hashedPassword
	user.UpdatedAt = time.Now()

	err = a.userRepo.UpdateUserRepo(ctx, user)
	if err != nil {
		return errors.New("修改密码失败")
	}

	return nil
}
