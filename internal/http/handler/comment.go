package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandlerNew struct {
	commentUsecase *usecase.CommentUseCase
	userUsecase    *usecase.UserUseCase
}

func NewCommentHandlerNew(commentUsecase *usecase.CommentUseCase, userUsecase *usecase.UserUseCase) *CommentHandlerNew {
	return &CommentHandlerNew{
		commentUsecase: commentUsecase,
		userUsecase:    userUsecase,
	}
}

// GetComments 获取评论列表
func (h *CommentHandlerNew) GetComments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	articleIdStr := c.Query("articleId")
	discussionIdStr := c.Query("discussionId")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 处理过滤参数
	var articleId, discussionId *int64
	if articleIdStr != "" {
		if id, err := strconv.ParseInt(articleIdStr, 10, 64); err == nil {
			articleId = &id
		}
	}
	if discussionIdStr != "" {
		if id, err := strconv.ParseInt(discussionIdStr, 10, 64); err == nil {
			discussionId = &id
		}
	}

	// 获取评论列表
	comments, total, err := h.commentUsecase.GetCommentsWithPagination(c.Request.Context(), page, pageSize, articleId, discussionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	// 提取并去重UserID
	userIDSet := make(map[int64]bool)
	for _, comment := range comments {
		userIDSet[comment.UserID] = true
	}

	// 批量查询用户信息
	userMap := make(map[int64]*entity.User)
	for userID := range userIDSet {
		user, err := h.userUsecase.GetUserById(c.Request.Context(), userID)
		if err == nil {
			userMap[userID] = user
		}
	}

	// 转换为响应格式
	commentResponses := make([]entity.CommentResponse, 0)
	for _, comment := range comments {
		user := userMap[comment.UserID]
		commentResponses = append(commentResponses, convertToCommentResponse(comment, user))
	}

	paginatedData := entity.NewPaginatedResponse(commentResponses, int(total), page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// CreateComment 创建评论
func (h *CommentHandlerNew) CreateComment(c *gin.Context) {
	var req entity.CommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "用户未登录"))
		return
	}

	// 创建评论实体
	comment := &entity.Comment{
		Content:      req.Content,
		ArticleID:    req.ArticleID,
		DiscussionID: req.DiscussionID,
		ParentID:     req.ParentID,
		UserID:       userID.(int64),
		Status:       "pending", // 默认状态为待审核
	}

	// 保存评论
	err := h.commentUsecase.AddComment(c.Request.Context(), comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
}

// UpdateComment 更新评论
func (h *CommentHandlerNew) UpdateComment(c *gin.Context) {
	idStr := c.Param("id")
	commentID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "评论ID格式错误"))
		return
	}

	var req entity.CommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "用户未登录"))
		return
	}

	// 获取原评论信息
	existingComment, err := h.commentUsecase.GetCommentById(c.Request.Context(), commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, entity.NewErrorResponse(404, "评论不存在"))
		return
	}

	// 检查权限：只有评论作者或管理员可以修改
	userRole, _ := c.Get("user_role")
	if existingComment.UserID != userID.(int64) && userRole != "admin" {
		c.JSON(http.StatusForbidden, entity.NewErrorResponse(403, "无权限修改此评论"))
		return
	}

	// 更新评论内容
	existingComment.Content = req.Content

	err = h.commentUsecase.UpdateComment(c.Request.Context(), existingComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "更新成功"))
}

// DeleteComment 删除评论
func (h *CommentHandlerNew) DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	commentID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "评论ID格式错误"))
		return
	}

	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "用户未登录"))
		return
	}

	// 获取原评论信息
	existingComment, err := h.commentUsecase.GetCommentById(c.Request.Context(), commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, entity.NewErrorResponse(404, "评论不存在"))
		return
	}

	// 检查权限：只有评论作者或管理员可以删除
	userRole, _ := c.Get("user_role")
	if existingComment.UserID != userID.(int64) && userRole != "admin" {
		c.JSON(http.StatusForbidden, entity.NewErrorResponse(403, "无权限删除此评论"))
		return
	}

	// 删除评论
	err = h.commentUsecase.DeleteComment(c.Request.Context(), commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
}

func convertToCommentResponse(comment *entity.Comment, user *entity.User) entity.CommentResponse {
	response := entity.CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		Status:    comment.Status,
		CreatedAt: comment.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: comment.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	if comment.ArticleID != nil {
		response.ArticleID = *comment.ArticleID
	}
	if comment.DiscussionID != nil {
		response.DiscussionID = *comment.DiscussionID
	}
	if comment.ParentID != nil {
		response.ParentID = *comment.ParentID
	}

	if user != nil {
		response.Author = entity.AuthorResponse{
			ID:       user.ID,
			Username: user.Username,
			Avatar:   user.Avatar,
		}
	} else {
		response.Author = entity.AuthorResponse{
			ID:       comment.UserID,
			Username: "未知用户",
			Avatar:   "",
		}
	}

	return response
}
