package usecase

import (
	"blog/internal/entity"
	"context"
	"errors"
	"strings"
)

type CommentUseCase struct {
	commentRepo CommentRepo
	postRepo    PostRepo
	userRepo    UserRepo
}

func NewCommentUseCase(commentRepo CommentRepo, postRepo PostRepo, userRepo UserRepo) *CommentUseCase {
	return &CommentUseCase{
		commentRepo: commentRepo,
		postRepo:    postRepo,
		userRepo:    userRepo,
	}
}

// Create creates a top-level comment or a single-level reply.
func (uc *CommentUseCase) Create(ctx context.Context, postID, userID string, req entity.CreateCommentRequest) (*entity.CommentResponse, error) {
	content := strings.TrimSpace(req.Content)
	if content == "" {
		return nil, errors.New("content cannot be empty")
	}
	if len([]rune(content)) > 1000 {
		return nil, errors.New("content exceeds 1000 characters")
	}
	if containsHTML(content) {
		return nil, errors.New("HTML content is not allowed")
	}

	// Ensure post exists (and is accessible). Here we only validate existence.
	if _, err := uc.postRepo.GetByID(ctx, postID); err != nil {
		return nil, errors.New("post not found")
	}

	var parentComment *entity.Comment
	if req.ParentID != nil && *req.ParentID != "" {
		pc, err := uc.commentRepo.GetByID(ctx, *req.ParentID)
		if err != nil {
			return nil, err
		}
		if pc.PostID != postID {
			return nil, errors.New("parent comment not in this post")
		}
		if pc.ParentID != nil {
			return nil, errors.New("only one-level replies are allowed")
		}
		parentComment = pc
	}

	comment := &entity.Comment{
		PostID:   postID,
		UserID:   userID,
		ParentID: req.ParentID,
		Content:  content,
	}

	if err := uc.commentRepo.Create(ctx, comment); err != nil {
		return nil, err
	}

	userInfo, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var replyToUser *entity.CommentUser
	if parentComment != nil {
		parentUser, err := uc.userRepo.GetByID(ctx, parentComment.UserID)
		if err == nil {
			replyToUser = &entity.CommentUser{
				Username: parentUser.Username,
				Avatar:   parentUser.Avatar,
			}
		}
	}

	return &entity.CommentResponse{
		ID:        comment.ID,
		ParentID:  comment.ParentID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		User: entity.CommentUser{
			Username: userInfo.Username,
			Avatar:   userInfo.Avatar,
		},
		ReplyToUser: replyToUser,
	}, nil
}

// List returns paginated top-level comments with optional one-level replies.
func (uc *CommentUseCase) List(ctx context.Context, postID string, page, limit int, order string, withReplies bool) (*entity.PaginatedCommentsResponse, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	topComments, total, err := uc.commentRepo.ListTopLevel(ctx, postID, page, limit, order)
	if err != nil {
		return nil, err
	}

	userCache := make(map[string]entity.CommentUser)
	getUser := func(uid string) (entity.CommentUser, error) {
		if u, ok := userCache[uid]; ok {
			return u, nil
		}
		user, err := uc.userRepo.GetByID(ctx, uid)
		if err != nil {
			return entity.CommentUser{}, err
		}
		u := entity.CommentUser{
			Username: user.Username,
			Avatar:   user.Avatar,
		}
		userCache[uid] = u
		return u, nil
	}

	// Build base responses
	responses := make([]entity.CommentResponse, 0, len(topComments))
	parentIDs := make([]string, 0, len(topComments))
	for _, c := range topComments {
		parentIDs = append(parentIDs, c.ID)
		user, err := getUser(c.UserID)
		if err != nil {
			continue
		}
		responses = append(responses, entity.CommentResponse{
			ID:        c.ID,
			ParentID:  c.ParentID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
			User:      user,
			Replies:   []entity.CommentResponse{},
		})
	}

	if withReplies && len(parentIDs) > 0 {
		replies, err := uc.commentRepo.ListReplies(ctx, postID, parentIDs, "asc")
		if err == nil {
			byParent := make(map[string][]entity.Comment)
			for _, r := range replies {
				if r.ParentID == nil {
					continue
				}
				byParent[*r.ParentID] = append(byParent[*r.ParentID], r)
			}

			// attach
			for i := range responses {
				rc := &responses[i]
				replyList := byParent[rc.ID]
				for _, r := range replyList {
					replyUser, err := getUser(r.UserID)
					if err != nil {
						continue
					}
					var replyTo *entity.CommentUser
					parentUser := rc.User
					replyTo = &entity.CommentUser{
						Username: parentUser.Username,
						Avatar:   parentUser.Avatar,
					}
					rc.Replies = append(rc.Replies, entity.CommentResponse{
						ID:          r.ID,
						ParentID:    r.ParentID,
						Content:     r.Content,
						CreatedAt:   r.CreatedAt,
						User:        replyUser,
						ReplyToUser: replyTo,
					})
				}
			}
		}
	}

	totalPages := (int(total) + limit - 1) / limit

	return &entity.PaginatedCommentsResponse{
		Data: responses,
		Pagination: entity.Pagination{
			Total:      int(total),
			Page:       page,
			Limit:      limit,
			TotalPages: totalPages,
		},
	}, nil
}

func containsHTML(content string) bool {
	// Lightweight guard: reject if HTML tag markers appear.
	return strings.Contains(content, "<") || strings.Contains(content, ">")
}
