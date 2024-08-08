package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MenuHandler struct {
	menuUseCase *usecase.MenuUseCase
}

func NewMenuHandler(u *usecase.MenuUseCase) *MenuHandler {
	return &MenuHandler{menuUseCase: u}
}

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (h *MenuHandler) AddMenu(c *gin.Context) {
	req := entity.MenuReq{}
	resp := new(ApiResponse)
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := h.menuUseCase.AddMenu(c, req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}
func (h *MenuHandler) Detail(c *gin.Context) {
	resp := new(ApiResponse)
	path := c.Param("path")
	result, err := h.menuUseCase.DetailByPath(c, path)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	resp.Data = result
	c.JSON(http.StatusOK, resp)
	return
}

func (h *MenuHandler) List(c *gin.Context) {
	resp := new(ApiResponse)
	menus, err := h.menuUseCase.List(c)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	resp.Data = menus
	c.JSON(http.StatusOK, resp)
	return
}

func (h *MenuHandler) UpdateMenu(c *gin.Context) {
	resp := new(ApiResponse)
	mid := c.Param("mid")
	menuId, err := strconv.Atoi(mid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	var req entity.MenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}
	req.ID = int64(menuId)
	if err := h.menuUseCase.UpdateMenu(c, req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
}

func (h *MenuHandler) DeleteMenuById(c *gin.Context) {
	resp := new(ApiResponse)
	mid := c.Param("mid")
	menuId, err := strconv.Atoi(mid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}
	if err := h.menuUseCase.DeleteMenu(c, int64(menuId)); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *MenuHandler) DeleteMenuBatch(c *gin.Context) {
	resp := new(ApiResponse)
	var req IdListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}
	if err := h.menuUseCase.DeleteMenusBatch(c, req.IDList); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}
