package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	tagUseCase *usecase.TagUseCase
}

func NewTagHandler(tagUseCase *usecase.TagUseCase) *TagHandler {
	return &TagHandler{tagUseCase: tagUseCase}
}

// ListTags - GET /tags
func (h *TagHandler) ListTags(c *gin.Context) {
	tags, err := h.tagUseCase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tags)
}

// CreateTag - POST /tags
func (h *TagHandler) CreateTag(c *gin.Context) {
	var req entity.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.tagUseCase.Create(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tag created successfully"})
}

// DeleteTag - DELETE /tags/:id
func (h *TagHandler) DeleteTag(c *gin.Context) {
	id := c.Param("id")

	if err := h.tagUseCase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
