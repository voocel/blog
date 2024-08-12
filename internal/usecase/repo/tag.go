package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

type TagRepo struct {
	db *gorm.DB
}

func NewTagRepo(db *gorm.DB) *TagRepo {
	return &TagRepo{db: db}
}

func (t TagRepo) AddTagRepo(ctx context.Context, tag *entity.Tag) error {
	return t.db.WithContext(ctx).Create(tag).Error
}

func (t TagRepo) AddTagsRepo(ctx context.Context, tags []*entity.Tag) error {
	return t.db.WithContext(ctx).Create(tags).Error
}

func (t TagRepo) GetTagByIdRepo(ctx context.Context, id int64) (*entity.Tag, error) {
	tag := new(entity.Tag)
	err := t.db.WithContext(ctx).Where("id = ?", id).First(tag).Error
	return tag, err
}

func (t TagRepo) GetTagByNameRepo(ctx context.Context, name string) (*entity.Tag, error) {
	tag := new(entity.Tag)
	err := t.db.WithContext(ctx).Where("name = ?", name).First(tag).Error
	return tag, err
}

func (t TagRepo) GetTagByNameExistRepo(ctx context.Context, name string) (bool, error) {
	t.db.WithContext(ctx).Where("name = ?", name).First(&entity.Tag{})
	return t.db.RowsAffected > 0, t.db.Error
}

func (t TagRepo) GetTagsRepo(ctx context.Context) ([]*entity.Tag, error) {
	var tags []*entity.Tag
	return tags, t.db.WithContext(ctx).Find(&tags).Error
}

func (t TagRepo) UpdateTagRepo(ctx context.Context, tag *entity.Tag) error {
	return t.db.WithContext(ctx).Updates(tag).Error
}

func (t TagRepo) DeleteTagRepo(ctx context.Context, id int64) error {
	return t.db.WithContext(ctx).Delete(&entity.Tag{}, id).Error
}

func (t TagRepo) DeleteTagsBatchRepo(ctx context.Context, ids []int64) error {
	return t.db.WithContext(ctx).Delete(&entity.Tag{}, ids).Error
}
