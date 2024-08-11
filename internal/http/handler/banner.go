package handler

import (
	"blog/config"
	"blog/internal/entity"
	"blog/internal/usecase"
	"blog/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type BannerHandler struct {
	bannerUseCase *usecase.BannerUseCase
}

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

func NewBannerHandler(u *usecase.BannerUseCase) *BannerHandler {
	return &BannerHandler{
		bannerUseCase: u,
	}
}

func (h *BannerHandler) Create(c *gin.Context) {
	resp := new(ApiResponse)
	var req entity.Banner
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

func (h *BannerHandler) CreateBanner(c *gin.Context) {
	resp := new(ApiResponse)
	form, err := c.MultipartForm()
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		resp.Code = 1
		resp.Message = "图片不存在"
		c.JSON(http.StatusOK, resp)
		return
	}

	date := time.Now().Format("2006-01-02")
	path := "static/banner/" + date + "/"

	var resData []FileUploadResponse
	for _, fileHeader := range fileList {
		filename := fileHeader.Filename
		filePathname := path + filename
		item := FileUploadResponse{
			IsSuccess: true,
			FileName:  filePathname,
		}
		if err = c.SaveUploadedFile(fileHeader, filePathname); err != nil {
			item.IsSuccess = false
			item.Msg = err.Error()
		} else {
			err = h.bannerUseCase.AddBanner(c, &entity.Banner{
				Name:      filename,
				Path:      config.Conf.App.Domain + "/" + filePathname,
				Hash:      "",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})
			if err != nil {
				log.Errorf("save to banner err: %v", err)
			}
		}
		resData = append(resData, item)
	}
	resp.Data = resData
	c.JSON(http.StatusOK, resp)
	return
}

func (h *BannerHandler) UpdateBanner(c *gin.Context) {
	resp := new(ApiResponse)
	var req entity.BannerUpdateReq
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
