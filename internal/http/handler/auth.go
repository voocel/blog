package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"blog/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
	userUseCase *usecase.UserUseCase
}

func NewAuthHandler(authUseCase *usecase.AuthUseCase, userUseCase *usecase.UserUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
		userUseCase: userUseCase,
	}
}

// Login - POST /auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req entity.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorw("Login bind JSON failed",
			log.Pair("error", err.Error()),
			log.Pair("ip", c.ClientIP()),
		)
		JSONError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	log.Infow("Login attempt",
		log.Pair("email", req.Email),
		log.Pair("ip", c.ClientIP()),
	)

	resp, err := h.authUseCase.Login(c.Request.Context(), req)
	if err != nil {
		log.Errorw("Login failed",
			log.Pair("email", req.Email),
			log.Pair("error", err.Error()),
			log.Pair("ip", c.ClientIP()),
		)
		JSONError(c, http.StatusUnauthorized, err.Error(), err)
		return
	}

	log.Infow("Login success",
		log.Pair("email", req.Email),
		log.Pair("username", resp.User.Username),
		log.Pair("ip", c.ClientIP()),
	)

	c.JSON(http.StatusOK, resp)
}

// Register - POST /auth/register
func (h *AuthHandler) Register(c *gin.Context) {
	var req entity.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorw("Register bind JSON failed",
			log.Pair("error", err.Error()),
			log.Pair("ip", c.ClientIP()),
		)
		JSONError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	log.Infow("Register attempt",
		log.Pair("email", req.Email),
		log.Pair("username", req.Username),
		log.Pair("ip", c.ClientIP()),
	)

	resp, err := h.authUseCase.Register(c.Request.Context(), req)
	if err != nil {
		log.Errorw("Register failed",
			log.Pair("email", req.Email),
			log.Pair("error", err.Error()),
			log.Pair("ip", c.ClientIP()),
		)
		JSONError(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	log.Infow("Register success",
		log.Pair("email", req.Email),
		log.Pair("username", resp.User.Username),
		log.Pair("ip", c.ClientIP()),
	)

	c.JSON(http.StatusCreated, resp)
}

// GetCurrentUser - GET /auth/me
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		log.Errorw("GetCurrentUser: user_id not found in context",
			log.Pair("ip", c.ClientIP()),
		)
		JSONError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		log.Errorw("GetCurrentUser: user_id is not a string",
			log.Pair("ip", c.ClientIP()),
		)
		JSONError(c, http.StatusInternalServerError, "Invalid user ID type", nil)
		return
	}

	user, err := h.authUseCase.GetCurrentUser(c.Request.Context(), userIDStr)
	if err != nil {
		log.Errorw("GetCurrentUser failed",
			log.Pair("user_id", userID),
			log.Pair("error", err.Error()),
			log.Pair("ip", c.ClientIP()),
		)
		JSONError(c, http.StatusNotFound, "User not found", err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// RefreshToken - POST /auth/refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req entity.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorw("RefreshToken bind JSON failed",
			log.Pair("error", err.Error()),
			log.Pair("ip", c.ClientIP()),
		)
		JSONError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	log.Infow("RefreshToken attempt",
		log.Pair("ip", c.ClientIP()),
	)

	resp, err := h.authUseCase.RefreshToken(c.Request.Context(), req)
	if err != nil {
		log.Errorw("RefreshToken failed",
			log.Pair("error", err.Error()),
			log.Pair("ip", c.ClientIP()),
		)
		JSONError(c, http.StatusUnauthorized, err.Error(), err)
		return
	}

	log.Infow("RefreshToken success",
		log.Pair("ip", c.ClientIP()),
	)

	c.JSON(http.StatusOK, resp)
}
