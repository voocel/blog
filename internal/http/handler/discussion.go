package handler

import (
	"blog/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DiscussionHandler struct {
}

func NewDiscussionHandler() *DiscussionHandler {
	return &DiscussionHandler{}
}

// GetDiscussions 获取讨论列表
func (h *DiscussionHandler) GetDiscussions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	// tag := c.Query("tag")
	// search := c.Query("search")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// todo
	var discussions []entity.DiscussionResponse
	discussion := entity.DiscussionResponse{
		ID:         1,
		Title:      "示例讨论",
		Content:    "这是一个示例讨论内容",
		Status:     "active",
		ViewCount:  10,
		ReplyCount: 5,
		Tags: []entity.TagResponse{
			{
				ID:    1,
				Name:  "技术",
				Title: "tech",
			},
		},
		Author: entity.AuthorResponse{
			ID:       1,
			Username: "admin",
			Avatar:   "/static/avatar/default.png",
		},
		CreatedAt: "2024-01-01T00:00:00Z",
		UpdatedAt: "2024-01-01T00:00:00Z",
	}
	discussions = append(discussions, discussion)

	paginatedData := entity.NewPaginatedResponse(discussions, 1, page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// GetDiscussion 获取讨论详情
func (h *DiscussionHandler) GetDiscussion(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "讨论ID格式错误"))
		return
	}

	// todo
	discussion := entity.DiscussionDetailResponse{
		ID:         1,
		Title:      "示例讨论",
		Content:    "这是一个示例讨论内容",
		Status:     "active",
		ViewCount:  10,
		ReplyCount: 2,
		Tags: []entity.TagResponse{
			{
				ID:    1,
				Name:  "技术",
				Title: "tech",
			},
		},
		Author: entity.AuthorResponse{
			ID:       1,
			Username: "admin",
			Avatar:   "/static/avatar/default.png",
		},
		Replies: []entity.ReplyResponse{
			{
				ID:      1,
				Content: "这是第一个回复",
				Author: entity.AuthorResponse{
					ID:       2,
					Username: "user1",
					Avatar:   "/static/avatar/default.png",
				},
				CreatedAt: "2024-01-01T01:00:00Z",
			},
			{
				ID:      2,
				Content: "这是第二个回复",
				Author: entity.AuthorResponse{
					ID:       3,
					Username: "user2",
					Avatar:   "/static/avatar/default.png",
				},
				CreatedAt: "2024-01-01T02:00:00Z",
			},
		},
		CreatedAt: "2024-01-01T00:00:00Z",
		UpdatedAt: "2024-01-01T00:00:00Z",
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(discussion, "获取成功"))
}

// CreateDiscussion 创建讨论
func (h *DiscussionHandler) CreateDiscussion(c *gin.Context) {
	var req entity.DiscussionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	// todo
	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
}

// UpdateDiscussion 更新讨论
func (h *DiscussionHandler) UpdateDiscussion(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "讨论ID格式错误"))
		return
	}

	var req entity.DiscussionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "更新成功"))
}

// DeleteDiscussion 删除讨论
func (h *DiscussionHandler) DeleteDiscussion(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "讨论ID格式错误"))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
}
