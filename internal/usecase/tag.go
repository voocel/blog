package usecase

import (
	"blog/internal/entity"
	"context"
)

type TagUseCase struct {
	tagRepo TagRepo
}

func NewTagUseCase(tagRepo TagRepo) *TagUseCase {
	return &TagUseCase{tagRepo: tagRepo}
}

func (uc *TagUseCase) Create(ctx context.Context, req entity.CreateTagRequest) error {
	tag := &entity.Tag{
		Name: req.Name,
	}
	return uc.tagRepo.Create(ctx, tag)
}

func (uc *TagUseCase) List(ctx context.Context) ([]entity.TagResponse, error) {
	tags, err := uc.tagRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]entity.TagResponse, len(tags))
	for i, tag := range tags {
		responses[i] = entity.TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		}
	}

	return responses, nil
}

func (uc *TagUseCase) Delete(ctx context.Context, id int64) error {
	return uc.tagRepo.Delete(ctx, id)
}
