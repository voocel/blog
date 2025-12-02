package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strconv"
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
	category := c.Query("category")
	search := c.Query("search")

	filters := map[string]interface{}{
		"status": "published",
		// Only show posts with publish date <= current date (scheduled publishing)
		"beforeDate": time.Now().Format("2006-01-02"),
	}
	if category != "" {
		filters["categoryId"] = category
	}
	if search != "" {
		filters["search"] = search
	}

	result, err := h.postUseCase.List(c.Request.Context(), filters, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// ListAllPosts - GET /admin/posts (Admin API)
func (h *PostHandler) ListAllPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	category := c.Query("category")
	status := c.Query("status") // all | published | draft
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetPost - GET /posts/:id (Public API)
func (h *PostHandler) GetPost(c *gin.Context) {
	id := c.Param("id")

	post, err := h.postUseCase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.Status != "published" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if publish date has arrived (scheduled publishing)
	currentDate := time.Now().Format("2006-01-02")
	if post.Date > currentDate {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetPostAdmin - GET /admin/posts/:id (Admin API)
func (h *PostHandler) GetPostAdmin(c *gin.Context) {
	id := c.Param("id")

	post, err := h.postUseCase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// CreatePost - POST /posts
func (h *PostHandler) CreatePost(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req entity.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"details": err.Error(),
		})
		return
	}

	if err := h.postUseCase.Create(c.Request.Context(), req, username.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

// UpdatePost - PUT /posts/:id
func (h *PostHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var req entity.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"details": err.Error(),
		})
		return
	}

	if err := h.postUseCase.Update(c.Request.Context(), id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	post, _ := h.postUseCase.GetByID(c.Request.Context(), id)
	c.JSON(http.StatusOK, post)
}

// DeletePost - DELETE /posts/:id
func (h *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")

	if err := h.postUseCase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
