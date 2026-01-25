package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

// UpdateProfile - PUT /users/profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		JSONError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		JSONError(c, http.StatusInternalServerError, "Invalid user ID type", nil)
		return
	}

	var req entity.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	if err := h.userUseCase.UpdateProfile(c.Request.Context(), userIDStr, req); err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	user, _ := h.userUseCase.GetByID(c.Request.Context(), userIDStr)
	c.JSON(http.StatusOK, user)
}

// ListUsersAdmin - GET /admin/users
func (h *UserHandler) ListUsersAdmin(c *gin.Context) {
	users, err := h.userUseCase.ListAll(c.Request.Context())
	if err != nil {
		JSONError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	c.JSON(http.StatusOK, users)
}

// UpdateUserStatus - PATCH /admin/users/:id/status
func (h *UserHandler) UpdateUserStatus(c *gin.Context) {
	userID := c.Param("id")
	var payload struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		JSONError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}
	status := strings.ToLower(payload.Status)
	updated, err := h.userUseCase.UpdateStatus(c.Request.Context(), userID, status)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			JSONError(c, http.StatusNotFound, "User not found", err)
			return
		}
		JSONError(c, http.StatusBadRequest, err.Error(), err)
		return
	}
	c.JSON(http.StatusOK, updated)
}
