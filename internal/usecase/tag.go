package usecase

import (
	"blog/internal/entity"
	"context"
)

type TagUseCase struct {
	repo TagRepo
}

func NewTagUseCase(repo TagRepo) *TagUseCase {
	return &TagUseCase{repo: repo}
}

func (t *TagUseCase) AddTag(ctx context.Context, req entity.TagReq) error {
	tag := new(entity.Tag)
	tag.Name = req.Name
	return t.repo.AddTagRepo(ctx, tag)
}

func (t *TagUseCase) GetTagById(ctx context.Context, id int64) (*entity.Tag, error) {
	return t.repo.GetTagByIdRepo(ctx, id)
}
func (t *TagUseCase) GetTagByName(ctx context.Context, name string) (*entity.Tag, error) {
	return t.repo.GetTagByNameRepo(ctx, name)
}

func (t *TagUseCase) GetTags(ctx context.Context) ([]*entity.Tag, error) {
	return t.repo.GetTagsRepo(ctx)
}

func (t *TagUseCase) UpdateTag(ctx context.Context, req entity.TagReq) error {
	tag := new(entity.Tag)
	tag.ID = req.ID
	tag.Name = req.Name
	return t.repo.UpdateTagRepo(ctx, tag)
}

func (t *TagUseCase) DeleteTag(ctx context.Context, id int64) error {
	return t.repo.DeleteTagRepo(ctx, id)
}

func (t *TagUseCase) DeleteTagBatch(ctx context.Context, ids []int64) error {
	return t.repo.DeleteTagsBatchRepo(ctx, ids)
}
