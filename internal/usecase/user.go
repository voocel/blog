package usecase

import (
	"blog/internal/entity"
	"context"
	"errors"
	"time"
)

type UserUseCase struct {
	userRepo UserRepo
}

func NewUserUseCase(userRepo UserRepo) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) GetByID(ctx context.Context, id string) (*entity.UserResponse, error) {
	user, err := uc.userRepo.GetByID(ctx, id)
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

func (uc *UserUseCase) UpdateProfile(ctx context.Context, id string, req entity.UpdateProfileRequest) error {
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	return uc.userRepo.Update(ctx, user)
}

// ListAll returns all users for admin management.
func (uc *UserUseCase) ListAll(ctx context.Context) ([]entity.AdminUserResponse, error) {
	users, err := uc.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]entity.AdminUserResponse, 0, len(users))
	for _, u := range users {
		resp = append(resp, entity.AdminUserResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			Role:     u.Role,
			Status:   u.Status,
			Provider: u.Provider,
			Avatar:   u.Avatar,
			JoinedAt: u.CreatedAt.Format(time.RFC3339),
		})
	}
	return resp, nil
}

// UpdateStatus updates user's status to active/banned.
func (uc *UserUseCase) UpdateStatus(ctx context.Context, id, status string) (*entity.AdminUserResponse, error) {
	if status != "active" && status != "banned" {
		return nil, errors.New("invalid status")
	}

	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user.Status != status {
		user.Status = status
		if err := uc.userRepo.Update(ctx, user); err != nil {
			return nil, err
		}
		// Revoke existing tokens on status change (ban/unban).
		if err := uc.userRepo.BumpTokenVersion(ctx, id); err != nil {
			return nil, err
		}
	}

	return &entity.AdminUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Status:   user.Status,
		Provider: user.Provider,
		Avatar:   user.Avatar,
		JoinedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}
