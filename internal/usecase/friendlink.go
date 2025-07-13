package usecase

import (
	"blog/internal/entity"
	"context"
	"errors"
)

type FriendlinkUseCase struct {
	repo FriendlinkRepo
}

func NewFriendlinkUseCase(repo FriendlinkRepo) *FriendlinkUseCase {
	return &FriendlinkUseCase{repo: repo}
}

func (uc *FriendlinkUseCase) GetFriendlinks(ctx context.Context, page, pageSize int, status string) ([]*entity.FriendlinkResponse, int64, error) {
	friendlinks, total, err := uc.repo.GetFriendlinksRepo(ctx, page, pageSize, status)
	if err != nil {
		return nil, 0, err
	}

	var responses []*entity.FriendlinkResponse
	for _, friendlink := range friendlinks {
		responses = append(responses, uc.convertToFriendlinkResponse(friendlink))
	}

	return responses, total, nil
}

func (uc *FriendlinkUseCase) CreateFriendlink(ctx context.Context, req *entity.FriendlinkRequest) error {
	friendlink := &entity.FriendLink{
		Name:        req.Name,
		URL:         req.URL,
		Logo:        req.Logo,
		Description: req.Description,
		Status:      req.Status,
		SortOrder:   req.SortOrder,
	}

	if friendlink.Status == "" {
		friendlink.Status = "active"
	}

	return uc.repo.AddFriendlinkRepo(ctx, friendlink)
}

// UpdateFriendlink 更新友链
func (uc *FriendlinkUseCase) UpdateFriendlink(ctx context.Context, id int64, req *entity.FriendlinkRequest) error {
	friendlink, err := uc.repo.GetFriendlinkByIdRepo(ctx, id)
	if err != nil {
		return errors.New("友链不存在")
	}

	friendlink.Name = req.Name
	friendlink.URL = req.URL
	friendlink.Logo = req.Logo
	friendlink.Description = req.Description
	friendlink.Status = req.Status
	friendlink.SortOrder = req.SortOrder

	return uc.repo.UpdateFriendlinkRepo(ctx, friendlink)
}

func (uc *FriendlinkUseCase) DeleteFriendlink(ctx context.Context, id int64) error {
	_, err := uc.repo.GetFriendlinkByIdRepo(ctx, id)
	if err != nil {
		return errors.New("友链不存在")
	}

	return uc.repo.DeleteFriendlinkRepo(ctx, id)
}

func (uc *FriendlinkUseCase) GetFriendlinkByID(ctx context.Context, id int64) (*entity.FriendlinkResponse, error) {
	friendlink, err := uc.repo.GetFriendlinkByIdRepo(ctx, id)
	if err != nil {
		return nil, errors.New("友链不存在")
	}

	return uc.convertToFriendlinkResponse(friendlink), nil
}

func (uc *FriendlinkUseCase) convertToFriendlinkResponse(friendlink *entity.FriendLink) *entity.FriendlinkResponse {
	return &entity.FriendlinkResponse{
		ID:          friendlink.ID,
		Name:        friendlink.Name,
		URL:         friendlink.URL,
		Logo:        friendlink.Logo,
		Description: friendlink.Description,
		Status:      friendlink.Status,
		SortOrder:   friendlink.SortOrder,
		CreatedAt:   friendlink.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   friendlink.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
