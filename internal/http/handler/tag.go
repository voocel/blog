package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TagHandlerNew 标签处理器
type TagHandlerNew struct {
	tagUsecase *usecase.TagUseCase
}

// NewTagHandlerNew 创建标签处理器
func NewTagHandlerNew(tagUsecase *usecase.TagUseCase) *TagHandlerNew {
	return &TagHandlerNew{
		tagUsecase: tagUsecase,
	}
}

// GetTags 获取标签列表
func (h *TagHandlerNew) GetTags(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	// search := c.Query("search") // 暂时未使用

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// todo
	tags, err := h.tagUsecase.GetTags(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	tagResponses := make([]entity.TagResponse, 0)
	for _, tag := range tags {
		tagResponses = append(tagResponses, convertToTagResponse(tag))
	}

	paginatedData := entity.NewPaginatedResponse(tagResponses, len(tagResponses), page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// CreateTag 创建标签
func (h *TagHandlerNew) CreateTag(c *gin.Context) {
	var req entity.TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	err := h.tagUsecase.AddTag(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
}

// UpdateTag 更新标签
func (h *TagHandlerNew) UpdateTag(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "标签ID格式错误"))
		return
	}

	var req entity.TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	err = h.tagUsecase.UpdateTag(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "更新成功"))
}

// DeleteTag 删除标签
func (h *TagHandlerNew) DeleteTag(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "标签ID格式错误"))
		return
	}

	err = h.tagUsecase.DeleteTag(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
}

// GetTag 获取单个标签
func (h *TagHandlerNew) GetTag(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "标签ID格式错误"))
		return
	}

	tag, err := h.tagUsecase.GetTagById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, entity.NewErrorResponse(404, "标签不存在"))
		return
	}

	response := convertToTagResponse(tag)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(response, "获取成功"))
}

func convertToTagResponse(tag *entity.Tag) entity.TagResponse {
	return entity.TagResponse{
		ID:           tag.ID,
		Name:         tag.Name,
		Title:        tag.Title,
		Description:  tag.Description,
		Color:        tag.Color,
		ArticleCount: tag.ArticleCount,
		CreatedAt:    tag.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:    tag.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
