package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"blog/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MenuHandler struct {
	menuUseCase       *usecase.MenuUseCase
	bannerUseCase     *usecase.BannerUseCase
	menuBannerUsecase *usecase.MenuBannerUseCase
}

func NewMenuHandler(u *usecase.MenuUseCase, b *usecase.BannerUseCase, mb *usecase.MenuBannerUseCase) *MenuHandler {
	return &MenuHandler{menuUseCase: u, bannerUseCase: b, menuBannerUsecase: mb}
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

	if h.menuUseCase.IsTitlePathExist(c, req.Title, req.Path) {
		resp.Code = 1
		resp.Message = "菜单名称或路径已存在"
		c.JSON(http.StatusOK, resp)
		return
	}

	menu, err := h.menuUseCase.AddMenu(c, req)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	// 菜单关联banner
	var menuBannerList []*entity.MenuBanner
	for _, sort := range req.ImageSortList {
		menuBannerList = append(menuBannerList, &entity.MenuBanner{
			MenuID:   menu.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	if err := h.menuBannerUsecase.AddMenuBannerBatch(c, menuBannerList); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}
func (h *MenuHandler) DetailByPath(c *gin.Context) {
	resp := new(ApiResponse)
	path := c.Query("path")
	menu, err := h.menuUseCase.DetailByPath(c, path)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	menuBanners, err := h.menuBannerUsecase.GetMenuBannerByMenuId(c, menu.ID)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	var banners = make([]entity.Banner, 0)
	for _, banner := range menuBanners {
		if menu.ID != banner.MenuID {
			continue
		}
		bannerInfo, err := h.bannerUseCase.Detail(c, banner.BannerID)
		if err != nil {
			log.Error(err)
			continue
		}
		banners = append(banners, entity.Banner{
			ID:   banner.BannerID,
			Path: bannerInfo.Path,
		})
	}
	type MenuResponse struct {
		entity.Menu
		Banners []entity.Banner `json:"banners"`
	}
	result := MenuResponse{
		Menu:    *menu,
		Banners: banners,
	}
	resp.Data = result
	c.JSON(http.StatusOK, resp)
	return
}

func (h *MenuHandler) DetailById(c *gin.Context) {
	resp := new(ApiResponse)
	mid := c.Param("mid")
	menuId, err := strconv.Atoi(mid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}
	menu, err := h.menuUseCase.Detail(c, int64(menuId))
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = menu
	c.JSON(http.StatusOK, resp)
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
