package handler

import (
	"blog/internal/entity"
	"blog/internal/http/middleware"
	"blog/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// Login 用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	var req entity.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	response, err := h.authUseCase.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(response, "登录成功"))
}

// Register 用户注册
func (h *AuthHandler) Register(c *gin.Context) {
	var req entity.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	response, err := h.authUseCase.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(response, "注册成功"))
}

// RefreshToken 刷新令牌
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req entity.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	response, err := h.authUseCase.RefreshToken(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(response, "令牌刷新成功"))
}

// Logout 用户登出
func (h *AuthHandler) Logout(c *gin.Context) {
	// 简单的登出响应，实际令牌失效由客户端处理
	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "登出成功"))
}

// GetProfile 获取当前用户信息
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, exists := middleware.GetCurrentUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "用户信息不存在"))
		return
	}

	response, err := h.authUseCase.GetProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, entity.NewErrorResponse(404, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(response, "获取成功"))
}

// UpdateProfile 更新用户资料
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	userID, exists := middleware.GetCurrentUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "用户信息不存在"))
		return
	}

	var req entity.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	response, err := h.authUseCase.UpdateProfile(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(response, "更新成功"))
}

// ChangePassword 修改密码
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	userID, exists := middleware.GetCurrentUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "用户信息不存在"))
		return
	}

	var req entity.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	err := h.authUseCase.ChangePassword(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "密码修改成功"))
}
