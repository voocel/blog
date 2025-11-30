package handler

import (
	"blog/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	mediaUseCase *usecase.MediaUseCase
}

func NewMediaHandler(mediaUseCase *usecase.MediaUseCase) *MediaHandler {
	return &MediaHandler{mediaUseCase: mediaUseCase}
}

// UploadFile - POST /upload
func (h *MediaHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// 获取基础 URL
	baseURL := "http://" + c.Request.Host

	media, err := h.mediaUseCase.Upload(c.Request.Context(), file, baseURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, media)
}

// ListFiles - GET /files
func (h *MediaHandler) ListFiles(c *gin.Context) {
	files, err := h.mediaUseCase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, files)
}

// DeleteFile - DELETE /files/:id
func (h *MediaHandler) DeleteFile(c *gin.Context) {
	id := c.Param("id")

	if err := h.mediaUseCase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
