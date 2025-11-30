package usecase

import (
	"blog/internal/entity"
	"context"
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
