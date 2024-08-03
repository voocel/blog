package handler

import (
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHandler struct {
	categoryUsecase *usecase.CategoryUseCase
}

func NewCategoryHandler(u *usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: u}
}

type CategoryResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (h *CategoryHandler) List(c *gin.Context) {
	resp := new(ApiResponse)
	categories, err := h.categoryUsecase.List(c)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	var categoryList = make([]CategoryResponse, 0)
	for _, category := range categories {
		categoryList = append(categoryList, CategoryResponse{
			Label: category.Name,
			Value: category.Name,
		})
	}
	resp.Data = categoryList
	c.JSON(http.StatusOK, resp)
	return
}
