package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdvertHandler struct {
	AdvertUseCase *usecase.AdvertUseCase
}

func NewAdvertHandler(u *usecase.AdvertUseCase) *AdvertHandler {
	return &AdvertHandler{AdvertUseCase: u}
}

func (h *AdvertHandler) Create(c *gin.Context) {
	req := entity.AdvertReq{}
	resp := new(ApiResponse)
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := h.AdvertUseCase.AddAdvert(c, &req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *AdvertHandler) Detail(c *gin.Context) {
	resp := new(ApiResponse)
	aid := c.Param("id")
	advertId, err := strconv.Atoi(aid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	advert, err := h.AdvertUseCase.Detail(c, int64(advertId))
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	resp.Data = advert
	c.JSON(http.StatusOK, resp)
	return
}

func (h *AdvertHandler) List(c *gin.Context) {
	resp := new(ApiResponse)
	result, err := h.AdvertUseCase.List(c)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	resp.Data = result
	c.JSON(http.StatusOK, resp)
	return
}

func (h *AdvertHandler) Update(c *gin.Context) {
	req := entity.AdvertReq{}
	resp := new(ApiResponse)
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := h.AdvertUseCase.UpdateAdvert(c, &req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *AdvertHandler) DeleteById(c *gin.Context) {
	resp := new(ApiResponse)
	aid := c.Param("id")
	advertId, err := strconv.Atoi(aid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	err = h.AdvertUseCase.DeleteAdvert(c, int64(advertId))
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *AdvertHandler) DeleteBatch(c *gin.Context) {
	resp := new(ApiResponse)
	var req IdListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	if err = h.AdvertUseCase.DeleteAdvertBatch(c, req.IDList); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}
