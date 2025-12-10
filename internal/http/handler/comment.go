package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateComment - POST /posts/:postId/comments
func (h *CommentHandler) CreateComment(c *gin.Context) {
	postID := c.Param("id")

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(string)

	var req entity.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	comment, err := h.commentUseCase.Create(c.Request.Context(), postID, userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}
