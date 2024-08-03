package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	articleUsecase *usecase.ArticleUseCase
}

func NewArticleHandler(u *usecase.ArticleUseCase) *ArticleHandler {
	return &ArticleHandler{
		articleUsecase: u,
	}
}

func (h *ArticleHandler) Create(c *gin.Context) {
	req := entity.ArticleReq{}
	resp := new(ApiResponse)
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}
	if err := h.articleUsecase.CreateArticle(c, req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *ArticleHandler) List(c *gin.Context) {
	resp := new(ApiResponse)
	articles, err := h.articleUsecase.GetList(c)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
	}
	resp.Data = articles
	c.JSON(http.StatusOK, resp)
	return
}

func (h *ArticleHandler) Detail(c *gin.Context) {
	resp := new(ApiResponse)
	aid := c.Param("aid")
	articleId, err := strconv.Atoi(aid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	article, err := h.articleUsecase.GetDetailById(c, articleId)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
	}
	resp.Data = article
	c.JSON(http.StatusOK, resp)
	return
}
