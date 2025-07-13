package repo

import (
	"blog/internal/entity"
	"context"

	"gorm.io/gorm"
)

// DiscussionRepo 讨论Repository
type DiscussionRepo struct {
	db *gorm.DB
}

// NewDiscussionRepo 创建讨论Repository
func NewDiscussionRepo(db *gorm.DB) *DiscussionRepo {
	return &DiscussionRepo{db: db}
}

// AddDiscussionRepo 添加讨论
func (r *DiscussionRepo) AddDiscussionRepo(ctx context.Context, discussion *entity.Discussion) error {
	return r.db.WithContext(ctx).Create(discussion).Error
}

// GetDiscussionByIdRepo 通过ID获取讨论
func (r *DiscussionRepo) GetDiscussionByIdRepo(ctx context.Context, id int64) (*entity.Discussion, error) {
	var discussion entity.Discussion
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&discussion).Error
	return &discussion, err
}

// GetDiscussionsRepo 获取讨论列表
func (r *DiscussionRepo) GetDiscussionsRepo(ctx context.Context, page, pageSize int, tagId *int64, search string) ([]*entity.Discussion, int64, error) {
	var discussions []*entity.Discussion
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Discussion{})

	// 搜索条件
	if search != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// 标签过滤 - 这里需要实现标签关联查询
	if tagId != nil && *tagId > 0 {
		// TODO: 实现标签过滤，需要联表查询
	}

	// 获取总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&discussions).Error

	return discussions, total, err
}

// UpdateDiscussionRepo 更新讨论
func (r *DiscussionRepo) UpdateDiscussionRepo(ctx context.Context, discussion *entity.Discussion) error {
	return r.db.WithContext(ctx).Model(discussion).Updates(discussion).Error
}

// DeleteDiscussionRepo 删除讨论
func (r *DiscussionRepo) DeleteDiscussionRepo(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.Discussion{}, id).Error
}

// IncrementViewCountRepo 增加访问量
func (r *DiscussionRepo) IncrementViewCountRepo(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Model(&entity.Discussion{}).Where("id = ?", id).
		Update("view_count", gorm.Expr("view_count + 1")).Error
}

// AddDiscussionTagsRepo 添加讨论标签关联
func (r *DiscussionRepo) AddDiscussionTagsRepo(ctx context.Context, discussionId int64, tagIds []int64) error {
	// 这里需要实现讨论标签关联表的插入
	// 暂时使用简单的JSON字段存储
	return nil
}

// DeleteDiscussionTagsRepo 删除讨论标签关联
func (r *DiscussionRepo) DeleteDiscussionTagsRepo(ctx context.Context, discussionId int64) error {
	// 这里需要实现讨论标签关联表的删除
	return nil
}

// GetDiscussionTagsRepo 获取讨论的标签
func (r *DiscussionRepo) GetDiscussionTagsRepo(ctx context.Context, discussionId int64) ([]int64, error) {
	// 这里需要实现获取讨论的标签ID列表
	return []int64{}, nil
}

// GetDiscussionsByTagIdRepo 通过标签ID获取讨论
func (r *DiscussionRepo) GetDiscussionsByTagIdRepo(ctx context.Context, tagId int64) ([]*entity.Discussion, error) {
	// 这里需要实现通过标签ID获取讨论列表
	return []*entity.Discussion{}, nil
}

// ReplyRepo 回复Repository
type ReplyRepo struct {
	db *gorm.DB
}

// NewReplyRepo 创建回复Repository
func NewReplyRepo(db *gorm.DB) *ReplyRepo {
	return &ReplyRepo{db: db}
}

// AddReplyRepo 添加回复
func (r *ReplyRepo) AddReplyRepo(ctx context.Context, reply *entity.Reply) error {
	return r.db.WithContext(ctx).Create(reply).Error
}

// GetReplyByIdRepo 通过ID获取回复
func (r *ReplyRepo) GetReplyByIdRepo(ctx context.Context, id int64) (*entity.Reply, error) {
	var reply entity.Reply
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&reply).Error
	return &reply, err
}

// GetRepliesByDiscussionIdRepo 通过讨论ID获取回复列表
func (r *ReplyRepo) GetRepliesByDiscussionIdRepo(ctx context.Context, discussionId int64) ([]*entity.Reply, error) {
	var replies []*entity.Reply
	err := r.db.WithContext(ctx).Where("discussion_id = ?", discussionId).
		Order("created_at ASC").Find(&replies).Error
	return replies, err
}

// UpdateReplyRepo 更新回复
func (r *ReplyRepo) UpdateReplyRepo(ctx context.Context, reply *entity.Reply) error {
	return r.db.WithContext(ctx).Model(reply).Updates(reply).Error
}

// DeleteReplyRepo 删除回复
func (r *ReplyRepo) DeleteReplyRepo(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.Reply{}, id).Error
}

// GetReplyCountByDiscussionIdRepo 获取讨论的回复数量
func (r *ReplyRepo) GetReplyCountByDiscussionIdRepo(ctx context.Context, discussionId int64) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Reply{}).Where("discussion_id = ?", discussionId).Count(&count).Error
	return count, err
}
