package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FriendlinkHandler struct {
	friendlinkUsecase *usecase.FriendlinkUseCase
}

func NewFriendlinkHandler(friendlinkUsecase *usecase.FriendlinkUseCase) *FriendlinkHandler {
	return &FriendlinkHandler{
		friendlinkUsecase: friendlinkUsecase,
	}
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

	friendlinks, total, err := h.friendlinkUsecase.GetFriendlinks(c.Request.Context(), page, pageSize, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	friendlinkResponses := make([]entity.FriendlinkResponse, 0)
	for _, friendlink := range friendlinks {
		friendlinkResponses = append(friendlinkResponses, *friendlink)
	}

	paginatedData := entity.NewPaginatedResponse(friendlinkResponses, int(total), page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// CreateFriendlink 创建友链
func (h *FriendlinkHandler) CreateFriendlink(c *gin.Context) {
	var req entity.FriendlinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	err := h.friendlinkUsecase.CreateFriendlink(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
}

// UpdateFriendlink 更新友链
func (h *FriendlinkHandler) UpdateFriendlink(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "友链ID格式错误"))
		return
	}

	var req entity.FriendlinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	err = h.friendlinkUsecase.UpdateFriendlink(c.Request.Context(), id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "更新成功"))
}

// DeleteFriendlink 删除友链
func (h *FriendlinkHandler) DeleteFriendlink(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "友链ID格式错误"))
		return
	}

	err = h.friendlinkUsecase.DeleteFriendlink(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
}

// GetFriendlink 获取单个友链
func (h *FriendlinkHandler) GetFriendlink(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "友链ID格式错误"))
		return
	}

	friendlink, err := h.friendlinkUsecase.GetFriendlinkByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, entity.NewErrorResponse(404, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(friendlink, "获取成功"))
}
