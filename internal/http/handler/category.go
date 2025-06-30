package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandlerNew struct {
	categoryUsecase *usecase.CategoryUseCase
}

func NewCategoryHandlerNew(categoryUsecase *usecase.CategoryUseCase) *CategoryHandlerNew {
	return &CategoryHandlerNew{
		categoryUsecase: categoryUsecase,
	}
}

// GetCategories 获取分类列表
func (h *CategoryHandlerNew) GetCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	// search := c.Query("search")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// todo
	categories, err := h.categoryUsecase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	var categoryResponses []entity.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, convertToCategoryResponse(category))
	}

	paginatedData := entity.NewPaginatedResponse(categoryResponses, len(categoryResponses), page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// CreateCategory 创建分类
func (h *CategoryHandlerNew) CreateCategory(c *gin.Context) {
	var req entity.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	// todo
	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
}

// UpdateCategory 更新分类
func (h *CategoryHandlerNew) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "分类ID格式错误"))
		return
	}

	var req entity.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "更新成功"))
}

// DeleteCategory 删除分类
func (h *CategoryHandlerNew) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "分类ID格式错误"))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
}

func convertToCategoryResponse(category *entity.Category) entity.CategoryResponse {
	response := entity.CategoryResponse{
		ID:           strconv.FormatInt(category.ID, 10),
		Name:         category.Name,
		Slug:         category.Slug,
		Description:  category.Description,
		ArticleCount: category.ArticleCount,
		CreatedAt:    category.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:    category.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	if category.ParentID != nil {
		response.ParentID = strconv.FormatInt(*category.ParentID, 10)
	}

	return response
}
