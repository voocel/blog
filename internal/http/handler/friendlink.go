package handler

import (
	"blog/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FriendlinkHandler struct {
}

func NewFriendlinkHandler() *FriendlinkHandler {
	return &FriendlinkHandler{}
}

func (h *FriendlinkHandler) GetFriendlinks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	status := c.Query("status")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// todo
	var friendlinks []entity.FriendlinkResponse
	if status == "" || status == "active" {
		friendlink := entity.FriendlinkResponse{
			ID:          "1",
			Name:        "示例网站",
			URL:         "https://example.com",
			Logo:        "https://example.com/logo.png",
			Description: "这是一个示例友链",
			Status:      "active",
			SortOrder:   1,
			CreatedAt:   "2024-01-01T00:00:00Z",
			UpdatedAt:   "2024-01-01T00:00:00Z",
		}
		friendlinks = append(friendlinks, friendlink)
	}

	paginatedData := entity.NewPaginatedResponse(friendlinks, len(friendlinks), page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// CreateFriendlink 创建友链
func (h *FriendlinkHandler) CreateFriendlink(c *gin.Context) {
	var req entity.FriendlinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	// todo
	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
}

// UpdateFriendlink 更新友链
func (h *FriendlinkHandler) UpdateFriendlink(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "友链ID格式错误"))
		return
	}

	var req entity.FriendlinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "更新成功"))
}

// DeleteFriendlink 删除友链
func (h *FriendlinkHandler) DeleteFriendlink(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "友链ID格式错误"))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
} 