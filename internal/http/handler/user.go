package handler

import (
	"blog/internal/entity"
	"blog/internal/http/middleware"
	"blog/internal/usecase"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
)

type UserHandler struct {
	userUsecase *usecase.UserUseCase
}

func NewUserHandler(userUsecase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

type ApiResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    interface{}        `json:"data,omitempty"`
	Paging  *PagingInformation `json:"paging,omitempty"`
}

type PagingInformation struct {
	Page     int   `json:"page"`      // 当前页码
	PageSize int   `json:"page_size"` // 每页的数量
	Total    int64 `json:"total"`     // 数据集中的总记录数
	Pages    int   `json:"pages"`     // 总页数
}

func (h *UserHandler) Login(c *gin.Context) {
	req := entity.UserLoginReq{}
	resp := new(ApiResponse)
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	user, err := h.userUsecase.UserLogin(ctx, req)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	at, rt, err := middleware.GenerateToken(user)
	if err != nil {
		resp.Code = 1
		resp.Message = "generate token fail"
		c.JSON(http.StatusOK, resp)
		return
	}
	result := map[string]interface{}{
		"id":            user.ID,
		"username":      user.Username,
		"nickname":      user.Nickname,
		"sex":           user.Sex,
		"avatar":        user.Avatar,
		"access_token":  at,
		"refresh_token": rt,
	}
	resp.Message = "login successfully"
	resp.Data = result
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) Register(c *gin.Context) {
	req := entity.UserRegisterReq{}
	resp := new(ApiResponse)
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}
	m := make(map[string]interface{})
	if err := copier.Copy(&m, req); err != nil {
		resp.Code = 1
		resp.Message = "params invalid."
		c.JSON(http.StatusOK, resp)
		return
	}
	//filter.Filters(m)
	if err := h.userUsecase.UserRegister(c, req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Message = "register successfully"
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) Info(c *gin.Context) {
	resp := new(ApiResponse)
	// Get userinfo from token
	user, exists := c.Get("jwt-user")
	if !exists {
		resp.Code = 1
		resp.Message = "user not exists"
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Message = "ok"
	resp.Data = user
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) Logout(c *gin.Context) {
	resp := new(ApiResponse)
	// Get userinfo from token
	_, exists := c.Get("jwt-user")
	if !exists {
		resp.Code = 1
		resp.Message = "user not exists"
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Message = "ok"
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) List(c *gin.Context) {
	resp := new(ApiResponse)
	users, err := h.userUsecase.UserList(c)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Message = "ok"
	resp.Data = users
	c.JSON(http.StatusOK, resp)
	return
}
