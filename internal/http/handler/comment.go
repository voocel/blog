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
}

func NewCommentHandlerNew(commentUsecase *usecase.CommentUseCase) *CommentHandlerNew {
	return &CommentHandlerNew{
		commentUsecase: commentUsecase,
	}
}

// GetComments 获取评论列表
func (h *CommentHandlerNew) GetComments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	// articleIdStr := c.Query("articleId")
	// discussionIdStr := c.Query("discussionId")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// todo
	comments, err := h.commentUsecase.GetComments(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	var commentResponses []entity.CommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, convertToCommentResponse(comment))
	}

	paginatedData := entity.NewPaginatedResponse(commentResponses, len(commentResponses), page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// CreateComment 创建评论
func (h *CommentHandlerNew) CreateComment(c *gin.Context) {
	var req entity.CommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	// todo
	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
}

// UpdateComment 更新评论
func (h *CommentHandlerNew) UpdateComment(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "评论ID格式错误"))
		return
	}

	var req entity.CommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "更新成功"))
}

// DeleteComment 删除评论
func (h *CommentHandlerNew) DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "评论ID格式错误"))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
}

func convertToCommentResponse(comment *entity.Comment) entity.CommentResponse {
	response := entity.CommentResponse{
		ID:        strconv.FormatInt(comment.ID, 10),
		Content:   comment.Content,
		Status:    comment.Status,
		CreatedAt: comment.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: comment.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	if comment.ArticleID != nil {
		response.ArticleID = strconv.FormatInt(*comment.ArticleID, 10)
	}
	if comment.DiscussionID != nil {
		response.DiscussionID = strconv.FormatInt(*comment.DiscussionID, 10)
	}
	if comment.ParentID != nil {
		response.ParentID = strconv.FormatInt(*comment.ParentID, 10)
	}

	// 如果有用户信息，则设置作者信息
	if comment.User != nil {
		response.Author = entity.AuthorResponse{
			ID:       strconv.FormatInt(comment.User.ID, 10),
			Username: comment.User.Username,
			Avatar:   comment.User.Avatar,
		}
	}

	return response
}
