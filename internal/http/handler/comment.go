package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"blog/pkg/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentUseCase *usecase.CommentUseCase
}

func NewCommentHandler(commentUseCase *usecase.CommentUseCase) *CommentHandler {
	return &CommentHandler{commentUseCase: commentUseCase}
}

// ListComments - GET /posts/:postId/comments
func (h *CommentHandler) ListComments(c *gin.Context) {
	postID := c.Param("id")
	if !util.IsValidUUID(postID) {
		JSONError(c, http.StatusBadRequest, "Invalid post id", nil)
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	order := strings.ToLower(c.DefaultQuery("order", "desc"))
	withReplies := true
	if v := c.Query("withReplies"); v != "" {
		if parsed, err := strconv.ParseBool(v); err == nil {
			withReplies = parsed
		}
	}

	resp, err := h.commentUseCase.List(c.Request.Context(), postID, page, limit, order, withReplies)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateComment - POST /posts/:postId/comments
func (h *CommentHandler) CreateComment(c *gin.Context) {
	postID := c.Param("id")
	if !util.IsValidUUID(postID) {
		JSONError(c, http.StatusBadRequest, "Invalid post id", nil)
		return
	}

	userIDVal, exists := c.Get("user_id")
	if !exists {
		JSONError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	userID := userIDVal.(string)

	var req entity.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	comment, err := h.commentUseCase.Create(c.Request.Context(), postID, userID, req)
	if err != nil {
		JSONError(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// ListAllCommentsAdmin - GET /admin/comments
func (h *CommentHandler) ListAllCommentsAdmin(c *gin.Context) {
	comments, err := h.commentUseCase.ListAllAdmin(c.Request.Context())
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	c.JSON(http.StatusOK, comments)
}

// DeleteCommentAdmin - DELETE /admin/comments/:id
func (h *CommentHandler) DeleteCommentAdmin(c *gin.Context) {
	id := c.Param("id")
	if err := h.commentUseCase.DeleteAdmin(c.Request.Context(), id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			JSONError(c, http.StatusNotFound, "Comment not found", err)
			return
		}
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	c.Status(http.StatusNoContent)
}
