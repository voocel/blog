package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUseCase *usecase.CategoryUseCase
}

func NewCategoryHandler(categoryUseCase *usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{categoryUseCase: categoryUseCase}
}

// ListCategories - GET /categories
func (h *CategoryHandler) ListCategories(c *gin.Context) {
	categories, err := h.categoryUseCase.List(c.Request.Context())
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, categories)
}

// CreateCategory - POST /categories
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req entity.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	if err := h.categoryUseCase.Create(c.Request.Context(), req); err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

// DeleteCategory - DELETE /categories/:id
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if err := h.categoryUseCase.Delete(c.Request.Context(), id); err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.Status(http.StatusNoContent)
}
