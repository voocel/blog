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

// Login authenticates user and returns access/refresh tokens
func (uc *AuthUseCase) Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error) {
	user, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if user.Status == "banned" {
		return nil, errors.New("user is banned")
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	tokenPair, err := jwt.GenerateTokenPair(user)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresIn:    tokenPair.ExpiresIn,
		User: entity.UserResponse{
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
			Avatar:   user.Avatar,
		},
	}, nil
}

// GetCurrentUser retrieves current user info
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

// Register creates new user account and returns tokens
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
		Role:     "visitor", // Default role
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	tokenPair, err := jwt.GenerateTokenPair(user)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresIn:    tokenPair.ExpiresIn,
		User: entity.UserResponse{
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
			Avatar:   user.Avatar,
		},
	}, nil
}

// RefreshToken generates new access token using refresh token
func (uc *AuthUseCase) RefreshToken(ctx context.Context, req entity.RefreshTokenRequest) (*entity.RefreshTokenResponse, error) {
	claims, err := jwt.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, errors.New("invalid or expired refresh token")
	}

	// Get user from database to ensure user still exists and get latest info
	user, err := uc.userRepo.GetByID(ctx, claims.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if user.Status == "banned" {
		return nil, errors.New("user is banned")
	}
	tokenTV := claims.TokenVersion
	if tokenTV <= 0 {
		tokenTV = 1
	}
	userTV := user.TokenVersion
	if userTV <= 0 {
		userTV = 1
	}
	if tokenTV != userTV {
		return nil, errors.New("token revoked")
	}

	// Generate new token pair
	tokenPair, err := jwt.GenerateTokenPair(user)
	if err != nil {
		return nil, err
	}

	return &entity.RefreshTokenResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresIn:    tokenPair.ExpiresIn,
	}, nil
}
