package handler

import (
	"blog/internal/usecase"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	mediaUseCase *usecase.MediaUseCase
}

func NewMediaHandler(mediaUseCase *usecase.MediaUseCase) *MediaHandler {
	return &MediaHandler{mediaUseCase: mediaUseCase}
}

// UploadFile - POST /upload?type=avatar|post
func (h *MediaHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		JSONError(c, http.StatusBadRequest, "No file uploaded", err)
		return
	}

	// Get upload type from query parameter: avatar | post | default: post
	uploadType := c.DefaultQuery("type", "post")
	if uploadType != "avatar" && uploadType != "post" {
		uploadType = "post"
	}

	baseURL := buildBaseURL(c)

	media, err := h.mediaUseCase.Upload(c.Request.Context(), file, baseURL, uploadType)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidArgument) {
			JSONError(c, http.StatusBadRequest, err.Error(), err)
			return
		}
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusCreated, media)
}

// UploadAvatar - POST /users/avatar (Dedicated avatar upload endpoint)
func (h *MediaHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		JSONError(c, http.StatusBadRequest, "No file uploaded", err)
		return
	}

	baseURL := buildBaseURL(c)

	media, err := h.mediaUseCase.Upload(c.Request.Context(), file, baseURL, "avatar")
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidArgument) {
			JSONError(c, http.StatusBadRequest, err.Error(), err)
			return
		}
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusCreated, media)
}

// ListFiles - GET /files
func (h *MediaHandler) ListFiles(c *gin.Context) {
	files, err := h.mediaUseCase.List(c.Request.Context())
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, files)
}

// DeleteFile - DELETE /files/:id
func (h *MediaHandler) DeleteFile(c *gin.Context) {
	id := c.Param("id")

	if err := h.mediaUseCase.Delete(c.Request.Context(), id); err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.Status(http.StatusNoContent)
}

func buildBaseURL(c *gin.Context) string {
	scheme := strings.TrimSpace(c.GetHeader("X-Forwarded-Proto"))
	if scheme != "" {
		if idx := strings.Index(scheme, ","); idx > -1 {
			scheme = strings.TrimSpace(scheme[:idx])
		}
	} else if c.Request.TLS != nil {
		scheme = "https"
	} else {
		scheme = "http"
	}
	return scheme + "://" + c.Request.Host
}
