package usecase

import (
	"blog/internal/entity"
	"context"
	"errors"
	"regexp"
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
func (uc *CommentUseCase) Create(ctx context.Context, postID, userID int64, req entity.CreateCommentRequest) (*entity.CommentResponse, error) {
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
	if req.ParentID != nil && *req.ParentID != 0 {
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
func (uc *CommentUseCase) List(ctx context.Context, postID int64, page, limit int, order string, withReplies bool) (*entity.PaginatedCommentsResponse, error) {
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

	userCache := make(map[int64]entity.CommentUser)
	getUser := func(uid int64) (entity.CommentUser, error) {
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
	parentIDs := make([]int64, 0, len(topComments))
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
			byParent := make(map[int64][]entity.Comment)
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

// ListAllAdmin returns all comments with user and post context for moderation.
func (uc *CommentUseCase) ListAllAdmin(ctx context.Context) ([]entity.AdminCommentResponse, error) {
	comments, err := uc.commentRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	// Batch load users and posts to avoid N+1 queries.
	userIDSet := make(map[int64]struct{})
	postIDSet := make(map[int64]struct{})
	for _, c := range comments {
		if c.UserID != 0 {
			userIDSet[c.UserID] = struct{}{}
		}
		if c.PostID != 0 {
			postIDSet[c.PostID] = struct{}{}
		}
	}
	userIDs := make([]int64, 0, len(userIDSet))
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}
	postIDs := make([]int64, 0, len(postIDSet))
	for id := range postIDSet {
		postIDs = append(postIDs, id)
	}

	users, err := uc.userRepo.GetByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	userByID := make(map[int64]entity.CommentUser, len(users))
	for i := range users {
		userByID[users[i].ID] = entity.CommentUser{Username: users[i].Username, Avatar: users[i].Avatar}
	}

	posts, err := uc.postRepo.GetByIDs(ctx, postIDs)
	if err != nil {
		return nil, err
	}
	postTitleByID := make(map[int64]string, len(posts))
	for i := range posts {
		postTitleByID[posts[i].ID] = posts[i].Title
	}

	resp := make([]entity.AdminCommentResponse, 0, len(comments))
	for _, c := range comments {
		user, ok := userByID[c.UserID]
		if !ok {
			continue
		}
		title := postTitleByID[c.PostID]
		if title == "" {
			continue
		}
		resp = append(resp, entity.AdminCommentResponse{
			ID:        c.ID,
			ParentID:  c.ParentID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
			PostID:    c.PostID,
			PostTitle: title,
			User:      user,
		})
	}

	return resp, nil
}

// DeleteAdmin deletes a comment and its direct replies.
func (uc *CommentUseCase) DeleteAdmin(ctx context.Context, id int64) error {
	// Ensure exists
	if _, err := uc.commentRepo.GetByID(ctx, id); err != nil {
		return err
	}
	return uc.commentRepo.DeleteCascade(ctx, id)
}

func containsHTML(content string) bool {
	return htmlTagPattern.MatchString(content)
}

var htmlTagPattern = regexp.MustCompile(`(?i)<\s*/?\s*[a-z][^>]*>`)
