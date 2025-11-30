package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postUseCase *usecase.PostUseCase
}

func NewPostHandler(postUseCase *usecase.PostUseCase) *PostHandler {
	return &PostHandler{postUseCase: postUseCase}
}

// ListPosts - GET /posts
func (h *PostHandler) ListPosts(c *gin.Context) {
	// 获取查询参数
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	category := c.Query("category")
	tag := c.Query("tag")
	status := c.DefaultQuery("status", "published")
	search := c.Query("search")

	// 构建过滤条件
	filters := make(map[string]interface{})
	if category != "" {
		filters["categoryId"] = category
	}
	if tag != "" {
		filters["tagId"] = tag
	}
	if status != "" {
		filters["status"] = status
	}
	if search != "" {
		filters["search"] = search
	}

	// 获取文章列表
	result, err := h.postUseCase.List(c.Request.Context(), filters, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetPost - GET /posts/:id
func (h *PostHandler) GetPost(c *gin.Context) {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.postUseCase.Update(c.Request.Context(), id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回更新后的文章
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
