package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BannerHandler struct {
	bannerUseCase *usecase.BannerUseCase
}

func NewBannerHandler(u *usecase.BannerUseCase) *BannerHandler {
	return &BannerHandler{
		bannerUseCase: u,
	}
}

func (h *BannerHandler) Create(c *gin.Context) {
	resp := new(ApiResponse)
	var req entity.BannerReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	err = h.bannerUseCase.AddBanner(c, &req)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *BannerHandler) Detail(c *gin.Context) {
	resp := new(ApiResponse)
	bid := c.Param("bid")
	bannerId, err := strconv.Atoi(bid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	banner, err := h.bannerUseCase.Detail(c, int64(bannerId))
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = banner
	c.JSON(http.StatusOK, resp)
	return
}

func (h *BannerHandler) List(c *gin.Context) {
	resp := new(ApiResponse)
	banners, err := h.bannerUseCase.List(c)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = banners
	c.JSON(http.StatusOK, resp)
	return
}

func (h *BannerHandler) Update(c *gin.Context) {
	resp := new(ApiResponse)
	var req entity.BannerReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	err = h.bannerUseCase.UpdateBanner(c, &req)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *BannerHandler) DeleteById(c *gin.Context) {
	resp := new(ApiResponse)
	bid := c.Param("bid")
	bannerId, err := strconv.Atoi(bid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	err = h.bannerUseCase.DeleteBanner(c, int64(bannerId))
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *BannerHandler) DeleteBatch(c *gin.Context) {
	resp := new(ApiResponse)
	var req IdListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	err = h.bannerUseCase.DeleteBannerBatch(c, req.IDList)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}
