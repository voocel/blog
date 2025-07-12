package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAdminHandler struct {
	userUsecase *usecase.UserUseCase
}

func NewUserAdminHandler(userUsecase *usecase.UserUseCase) *UserAdminHandler {
	return &UserAdminHandler{
		userUsecase: userUsecase,
	}
}

// GetUsers 获取用户列表
func (h *UserAdminHandler) GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	search := c.Query("search")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	users, total, err := h.userUsecase.UserList(c.Request.Context(), page, pageSize, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	var userResponses []entity.UserAdminResponse
	for _, user := range users {
		userResponses = append(userResponses, convertToUserAdminResponse(user))
	}

	paginatedData := entity.NewPaginatedResponse(userResponses, int(total), page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// CreateUser 创建用户
func (h *UserAdminHandler) CreateUser(c *gin.Context) {
	var req entity.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	// todo
	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
}

// UpdateUser 更新用户
func (h *UserAdminHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "用户ID格式错误"))
		return
	}

	var req entity.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	// todo
	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "更新成功"))
}

// DeleteUser 删除用户
func (h *UserAdminHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	if _, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "用户ID格式错误"))
		return
	}

	// todo
	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
}

func convertToUserAdminResponse(user *entity.User) entity.UserAdminResponse {
	return entity.UserAdminResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Role:        user.Role,
		Status:      user.Status,
		Nickname:    user.Nickname,
		Website:     user.Website,
		Description: user.Description,
		CreatedAt:   user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
