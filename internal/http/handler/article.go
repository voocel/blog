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

type IdListReq struct {
	IDList []int64 `json:"id_list"`
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
		return
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

	article, err := h.articleUsecase.GetDetailById(c, int64(articleId))
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = article
	c.JSON(http.StatusOK, resp)
	return
}

func (h *ArticleHandler) DeleteArticleById(c *gin.Context) {
	resp := new(ApiResponse)
	aid := c.Param("aid")
	articleId, err := strconv.Atoi(aid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := h.articleUsecase.DeleteArticle(c, int64(articleId)); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *ArticleHandler) DeleteArticlesBatch(c *gin.Context) {
	resp := new(ApiResponse)
	var req IdListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := h.articleUsecase.DeleteArticles(c, req.IDList); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}
