package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DiscussionHandler struct {
	discussionUsecase *usecase.DiscussionUseCase
}

func NewDiscussionHandler(discussionUsecase *usecase.DiscussionUseCase) *DiscussionHandler {
	return &DiscussionHandler{
		discussionUsecase: discussionUsecase,
	}
}

// GetDiscussions 获取讨论列表
func (h *DiscussionHandler) GetDiscussions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	tagStr := c.Query("tag")
	search := c.Query("search")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	var tagId *int64
	if tagStr != "" {
		if tid, err := strconv.ParseInt(tagStr, 10, 64); err == nil {
			tagId = &tid
		}
	}

	discussions, total, err := h.discussionUsecase.GetDiscussions(c.Request.Context(), page, pageSize, tagId, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	// 转换为响应格式
	discussionResponses := make([]entity.DiscussionResponse, 0)
	for _, discussion := range discussions {
		discussionResponses = append(discussionResponses, *discussion)
	}

	paginatedData := entity.NewPaginatedResponse(discussionResponses, int(total), page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// GetDiscussion 获取讨论详情
func (h *DiscussionHandler) GetDiscussion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "讨论ID格式错误"))
		return
	}

	discussion, err := h.discussionUsecase.GetDiscussionDetail(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, entity.NewErrorResponse(404, err.Error()))
		return
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

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "用户未登录"))
		return
	}

	uid, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, "用户ID格式错误"))
		return
	}

	err := h.discussionUsecase.CreateDiscussion(c.Request.Context(), uid, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
}

// UpdateDiscussion 更新讨论
func (h *DiscussionHandler) UpdateDiscussion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "讨论ID格式错误"))
		return
	}

	var req entity.DiscussionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "用户未登录"))
		return
	}

	uid, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, "用户ID格式错误"))
		return
	}

	err = h.discussionUsecase.UpdateDiscussion(c.Request.Context(), id, uid, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "更新成功"))
}

// DeleteDiscussion 删除讨论
func (h *DiscussionHandler) DeleteDiscussion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "讨论ID格式错误"))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "用户未登录"))
		return
	}

	uid, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, "用户ID格式错误"))
		return
	}

	err = h.discussionUsecase.DeleteDiscussion(c.Request.Context(), id, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
}

// CreateReply 创建回复
func (h *DiscussionHandler) CreateReply(c *gin.Context) {
	idStr := c.Param("id")
	discussionID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "讨论ID格式错误"))
		return
	}

	var req struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "用户未登录"))
		return
	}

	uid, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, "用户ID格式错误"))
		return
	}

	err = h.discussionUsecase.CreateReply(c.Request.Context(), discussionID, uid, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "回复成功"))
}
