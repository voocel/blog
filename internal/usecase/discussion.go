package usecase

import (
	"blog/internal/entity"
	"context"
	"errors"
	"strconv"
	"strings"
)

// DiscussionUseCase 讨论UseCase
type DiscussionUseCase struct {
	discussionRepo DiscussionRepo
	replyRepo      ReplyRepo
	userRepo       UserRepo
	tagRepo        TagRepo
}

// NewDiscussionUseCase 创建讨论UseCase
func NewDiscussionUseCase(discussionRepo DiscussionRepo, replyRepo ReplyRepo, userRepo UserRepo, tagRepo TagRepo) *DiscussionUseCase {
	return &DiscussionUseCase{
		discussionRepo: discussionRepo,
		replyRepo:      replyRepo,
		userRepo:       userRepo,
		tagRepo:        tagRepo,
	}
}

// GetDiscussions 获取讨论列表
func (uc *DiscussionUseCase) GetDiscussions(ctx context.Context, page, pageSize int, tagId *int64, search string) ([]*entity.DiscussionResponse, int64, error) {
	discussions, total, err := uc.discussionRepo.GetDiscussionsRepo(ctx, page, pageSize, tagId, search)
	if err != nil {
		return nil, 0, err
	}

	var responses []*entity.DiscussionResponse
	for _, discussion := range discussions {
		response, err := uc.convertToDiscussionResponse(ctx, discussion)
		if err != nil {
			continue // 跳过转换失败的记录
		}
		responses = append(responses, response)
	}

	return responses, total, nil
}

// GetDiscussionDetail 获取讨论详情
func (uc *DiscussionUseCase) GetDiscussionDetail(ctx context.Context, id int64) (*entity.DiscussionDetailResponse, error) {
	// 增加访问量
	_ = uc.discussionRepo.IncrementViewCountRepo(ctx, id)

	discussion, err := uc.discussionRepo.GetDiscussionByIdRepo(ctx, id)
	if err != nil {
		return nil, errors.New("讨论不存在")
	}

	// 获取作者信息
	author, err := uc.userRepo.GetUserByIdRepo(ctx, discussion.UserID)
	if err != nil {
		return nil, err
	}

	// 获取标签信息
	tags, err := uc.getDiscussionTags(ctx, discussion)
	if err != nil {
		return nil, err
	}

	// 获取回复列表
	replies, err := uc.replyRepo.GetRepliesByDiscussionIdRepo(ctx, id)
	if err != nil {
		return nil, err
	}

	// 转换回复响应
	var replyResponses []entity.ReplyResponse
	for _, reply := range replies {
		replyAuthor, err := uc.userRepo.GetUserByIdRepo(ctx, reply.UserID)
		if err != nil {
			continue // 跳过获取作者失败的回复
		}

		replyResponses = append(replyResponses, entity.ReplyResponse{
			ID:      reply.ID,
			Content: reply.Content,
			Author: entity.AuthorResponse{
				ID:       replyAuthor.ID,
				Username: replyAuthor.Username,
				Avatar:   replyAuthor.Avatar,
			},
			CreatedAt: reply.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &entity.DiscussionDetailResponse{
		ID:         discussion.ID,
		Title:      discussion.Title,
		Content:    discussion.Content,
		Status:     discussion.Status,
		ViewCount:  discussion.ViewCount,
		ReplyCount: len(replyResponses),
		Tags:       tags,
		Author: entity.AuthorResponse{
			ID:       author.ID,
			Username: author.Username,
			Avatar:   author.Avatar,
		},
		Replies:   replyResponses,
		CreatedAt: discussion.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: discussion.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

// CreateDiscussion 创建讨论
func (uc *DiscussionUseCase) CreateDiscussion(ctx context.Context, userID int64, req *entity.DiscussionRequest) error {
	// 构建标签数组字符串
	var tagStr string
	if len(req.TagIds) > 0 {
		tagStrs := make([]string, len(req.TagIds))
		for i, tagId := range req.TagIds {
			tagStrs[i] = strconv.FormatInt(tagId, 10)
		}
		tagStr = strings.Join(tagStrs, ",")
	}

	discussion := &entity.Discussion{
		Title:   req.Title,
		Content: req.Content,
		Status:  req.Status,
		UserID:  userID,
		Tags:    entity.Array(strings.Split(tagStr, ",")),
	}

	if discussion.Status == "" {
		discussion.Status = "active"
	}

	return uc.discussionRepo.AddDiscussionRepo(ctx, discussion)
}

// UpdateDiscussion 更新讨论
func (uc *DiscussionUseCase) UpdateDiscussion(ctx context.Context, id, userID int64, req *entity.DiscussionRequest) error {
	discussion, err := uc.discussionRepo.GetDiscussionByIdRepo(ctx, id)
	if err != nil {
		return errors.New("讨论不存在")
	}

	// 检查权限（只有作者可以更新）
	if discussion.UserID != userID {
		return errors.New("无权限更新此讨论")
	}

	// 更新字段
	discussion.Title = req.Title
	discussion.Content = req.Content
	discussion.Status = req.Status

	// 更新标签
	if len(req.TagIds) > 0 {
		tagStrs := make([]string, len(req.TagIds))
		for i, tagId := range req.TagIds {
			tagStrs[i] = strconv.FormatInt(tagId, 10)
		}
		discussion.Tags = entity.Array(tagStrs)
	}

	return uc.discussionRepo.UpdateDiscussionRepo(ctx, discussion)
}

// DeleteDiscussion 删除讨论
func (uc *DiscussionUseCase) DeleteDiscussion(ctx context.Context, id, userID int64) error {
	discussion, err := uc.discussionRepo.GetDiscussionByIdRepo(ctx, id)
	if err != nil {
		return errors.New("讨论不存在")
	}

	// 检查权限（只有作者可以删除）
	if discussion.UserID != userID {
		return errors.New("无权限删除此讨论")
	}

	return uc.discussionRepo.DeleteDiscussionRepo(ctx, id)
}

// CreateReply 创建回复
func (uc *DiscussionUseCase) CreateReply(ctx context.Context, discussionID, userID int64, content string) error {
	_, err := uc.discussionRepo.GetDiscussionByIdRepo(ctx, discussionID)
	if err != nil {
		return errors.New("讨论不存在")
	}

	reply := &entity.Reply{
		Content:      content,
		DiscussionID: discussionID,
		UserID:       userID,
	}

	return uc.replyRepo.AddReplyRepo(ctx, reply)
}

// convertToDiscussionResponse 转换为讨论响应
func (uc *DiscussionUseCase) convertToDiscussionResponse(ctx context.Context, discussion *entity.Discussion) (*entity.DiscussionResponse, error) {
	author, err := uc.userRepo.GetUserByIdRepo(ctx, discussion.UserID)
	if err != nil {
		return nil, err
	}

	tags, err := uc.getDiscussionTags(ctx, discussion)
	if err != nil {
		return nil, err
	}

	replyCount, _ := uc.replyRepo.GetReplyCountByDiscussionIdRepo(ctx, discussion.ID)

	return &entity.DiscussionResponse{
		ID:         discussion.ID,
		Title:      discussion.Title,
		Content:    discussion.Content,
		Status:     discussion.Status,
		ViewCount:  discussion.ViewCount,
		ReplyCount: int(replyCount),
		Tags:       tags,
		Author: entity.AuthorResponse{
			ID:       author.ID,
			Username: author.Username,
			Avatar:   author.Avatar,
		},
		CreatedAt: discussion.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: discussion.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

// getDiscussionTags 获取讨论的标签信息
func (uc *DiscussionUseCase) getDiscussionTags(ctx context.Context, discussion *entity.Discussion) ([]entity.TagResponse, error) {
	var tags []entity.TagResponse

	if len(discussion.Tags) > 0 {
		for _, tagStr := range discussion.Tags {
			if tagStr == "" {
				continue
			}
			tagId, err := strconv.ParseInt(tagStr, 10, 64)
			if err != nil {
				continue
			}

			tag, err := uc.tagRepo.GetTagByIdRepo(ctx, tagId)
			if err != nil {
				continue
			}

			tags = append(tags, entity.TagResponse{
				ID:    tag.ID,
				Name:  tag.Name,
				Title: tag.Title,
			})
		}
	}

	return tags, nil
}
