package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeHandler struct {
	likeUseCase *usecase.LikeUseCase
}

func NewLikeHandler(likeUseCase *usecase.LikeUseCase) *LikeHandler {
	return &LikeHandler{likeUseCase: likeUseCase}
}

// GetLikes - GET /likes?slug=home
func (h *LikeHandler) GetLikes(c *gin.Context) {
	slug := c.Query("slug")
	if slug == "" {
		JSONError(c, http.StatusBadRequest, "Missing slug parameter", nil)
		return
	}

	count, err := h.likeUseCase.GetCount(c.Request.Context(), slug)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Failed to get like count", err)
		return
	}

	c.JSON(http.StatusOK, entity.LikeResponse{Count: count})
}

// Like - POST /likes?slug=home
func (h *LikeHandler) Like(c *gin.Context) {
	slug := c.Query("slug")
	if slug == "" {
		JSONError(c, http.StatusBadRequest, "Missing slug parameter", nil)
		return
	}

	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	count, err := h.likeUseCase.Like(c.Request.Context(), slug, ip, userAgent)
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Failed to like", err)
		return
	}

	c.JSON(http.StatusOK, entity.LikeResponse{Count: count})
}
