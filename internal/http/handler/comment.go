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
	postUseCase    *usecase.PostUseCase
}

func NewCommentHandler(commentUseCase *usecase.CommentUseCase, postUseCase *usecase.PostUseCase) *CommentHandler {
	return &CommentHandler{commentUseCase: commentUseCase, postUseCase: postUseCase}
}

// ListComments - GET /posts/:slug/comments
func (h *CommentHandler) ListComments(c *gin.Context) {
	slug := c.Param("slug")
	if !util.IsValidSlug(slug) {
		JSONError(c, http.StatusBadRequest, "Invalid post slug", nil)
		return
	}

	// Look up post by slug to get its ID
	post, err := h.postUseCase.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		JSONError(c, http.StatusNotFound, "Post not found", err)
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

	resp, err := h.commentUseCase.List(c.Request.Context(), post.ID, page, limit, order, withReplies)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateComment - POST /posts/:slug/comments
func (h *CommentHandler) CreateComment(c *gin.Context) {
	slug := c.Param("slug")
	if !util.IsValidSlug(slug) {
		JSONError(c, http.StatusBadRequest, "Invalid post slug", nil)
		return
	}

	// Look up post by slug to get its ID
	post, err := h.postUseCase.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		JSONError(c, http.StatusNotFound, "Post not found", err)
		return
	}

	userIDVal, exists := c.Get("user_id")
	if !exists {
		JSONError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	userID, ok := userIDVal.(int64)
	if !ok {
		JSONError(c, http.StatusInternalServerError, "Invalid user ID type", nil)
		return
	}

	var req entity.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	comment, err := h.commentUseCase.Create(c.Request.Context(), post.ID, userID, req)
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
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		JSONError(c, http.StatusBadRequest, "Invalid comment id", nil)
		return
	}
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
