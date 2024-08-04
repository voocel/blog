package handler

import (
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type StarHandler struct {
	starUsecase *usecase.StarUseCase
}

func NewStarHandler(u *usecase.StarUseCase) *StarHandler {
	return &StarHandler{starUsecase: u}
}

func (h *StarHandler) List(c *gin.Context) {
	resp := new(ApiResponse)
	result, err := h.starUsecase.GetStarsByUidRepo(c, 1)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = result
	c.JSON(http.StatusOK, resp)
	return
}

func (h *StarHandler) AddStar(c *gin.Context) {
	resp := new(ApiResponse)

	aid := c.Param("aid")
	articleId, err := strconv.Atoi(aid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := h.starUsecase.AddStar(c, 1, int64(articleId)); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *StarHandler) RemoveStar(c *gin.Context) {
	resp := new(ApiResponse)

	aid := c.Param("aid")
	articleId, err := strconv.Atoi(aid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := h.starUsecase.DeleteStar(c, 1, int64(articleId)); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}
