package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postUseCase *usecase.PostUseCase
}

func NewPostHandler(postUseCase *usecase.PostUseCase) *PostHandler {
	return &PostHandler{postUseCase: postUseCase}
}

// ListPublishedPosts - GET /posts (Public API)
func (h *PostHandler) ListPublishedPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	limit = clampLimit(limit, 100)
	category := c.Query("category")
	search := c.Query("search")

	filters := map[string]interface{}{
		"status": "published",
		// Only show posts where publish_at <= now (scheduled publishing).
		"beforePublishAt": time.Now(),
	}
	if category != "" {
		filters["categoryId"] = category
	}
	if search != "" {
		filters["search"] = search
	}

	result, err := h.postUseCase.List(c.Request.Context(), filters, page, limit)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	// Log homepage visit (only for first page without filters)
	if page <= 1 && category == "" && search == "" {
		go h.postUseCase.LogHomeVisit(c.ClientIP(), c.GetHeader("User-Agent"))
	}

	c.JSON(http.StatusOK, result)
}

// ListAllPosts - GET /admin/posts (Admin API)
func (h *PostHandler) ListAllPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	limit = clampLimit(limit, 100)
	category := c.Query("category")
	status := strings.ToLower(strings.TrimSpace(c.Query("status"))) // all | published | draft
	search := c.Query("search")

	filters := make(map[string]interface{})
	if category != "" {
		filters["categoryId"] = category
	}

	if status != "" && status != "all" {
		filters["status"] = status
	}
	if search != "" {
		filters["search"] = search
	}

	result, err := h.postUseCase.List(c.Request.Context(), filters, page, limit)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetPost - GET /posts/:id (Public API)
func (h *PostHandler) GetPost(c *gin.Context) {
	id := c.Param("id")
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	post, err := h.postUseCase.GetByIDWithAnalytics(c.Request.Context(), id, ip, userAgent)
	if err != nil {
		JSONError(c, http.StatusNotFound, "Post not found", err)
		return
	}

	if post.Status != "published" {
		JSONError(c, http.StatusNotFound, "Post not found", nil)
		return
	}

	// Check if publish date has arrived (scheduled publishing)
	if post.PublishAt.After(time.Now()) {
		JSONError(c, http.StatusNotFound, "Post not found", nil)
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetPostAdmin - GET /admin/posts/:id (Admin API)
func (h *PostHandler) GetPostAdmin(c *gin.Context) {
	id := c.Param("id")

	post, err := h.postUseCase.GetByID(c.Request.Context(), id)
	if err != nil {
		JSONError(c, http.StatusNotFound, "Post not found", err)
		return
	}

	c.JSON(http.StatusOK, post)
}

// CreatePost - POST /posts
func (h *PostHandler) CreatePost(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		JSONError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var req entity.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	if err := h.postUseCase.Create(c.Request.Context(), req, username.(string)); err != nil {
		if errors.Is(err, usecase.ErrInvalidArgument) {
			JSONError(c, http.StatusBadRequest, err.Error(), err)
			return
		}
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

// UpdatePost - PUT /posts/:id
func (h *PostHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var req entity.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	if err := h.postUseCase.Update(c.Request.Context(), id, req); err != nil {
		if errors.Is(err, usecase.ErrInvalidArgument) {
			JSONError(c, http.StatusBadRequest, err.Error(), err)
			return
		}
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	post, err := h.postUseCase.GetByID(c.Request.Context(), id)
	if err != nil {
		JSONError(c, http.StatusNotFound, "Post not found", err)
		return
	}
	c.JSON(http.StatusOK, post)
}

// DeletePost - DELETE /posts/:id
func (h *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")

	if err := h.postUseCase.Delete(c.Request.Context(), id); err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.Status(http.StatusNoContent)
}

func clampLimit(limit int, max int) int {
	if limit > max {
		return max
	}
	return limit
}
