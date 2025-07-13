package repo

import (
	"blog/internal/entity"
	"context"

	"gorm.io/gorm"
)

// FriendlinkRepo 友链Repository
type FriendlinkRepo struct {
	db *gorm.DB
}

// NewFriendlinkRepo 创建友链Repository
func NewFriendlinkRepo(db *gorm.DB) *FriendlinkRepo {
	return &FriendlinkRepo{db: db}
}

// AddFriendlinkRepo 添加友链
func (r *FriendlinkRepo) AddFriendlinkRepo(ctx context.Context, friendlink *entity.FriendLink) error {
	return r.db.WithContext(ctx).Create(friendlink).Error
}

// GetFriendlinkByIdRepo 通过ID获取友链
func (r *FriendlinkRepo) GetFriendlinkByIdRepo(ctx context.Context, id int64) (*entity.FriendLink, error) {
	var friendlink entity.FriendLink
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&friendlink).Error
	return &friendlink, err
}

// GetFriendlinksRepo 获取友链列表
func (r *FriendlinkRepo) GetFriendlinksRepo(ctx context.Context, page, pageSize int, status string) ([]*entity.FriendLink, int64, error) {
	var friendlinks []*entity.FriendLink
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.FriendLink{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("sort_order ASC, created_at DESC").Find(&friendlinks).Error

	return friendlinks, total, err
}

// UpdateFriendlinkRepo 更新友链
func (r *FriendlinkRepo) UpdateFriendlinkRepo(ctx context.Context, friendlink *entity.FriendLink) error {
	return r.db.WithContext(ctx).Model(friendlink).Updates(friendlink).Error
}

// DeleteFriendlinkRepo 删除友链
func (r *FriendlinkRepo) DeleteFriendlinkRepo(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.FriendLink{}, id).Error
}
