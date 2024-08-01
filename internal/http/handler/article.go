package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
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
}

func (h *ArticleHandler) List(c *gin.Context) {

}

func (h *ArticleHandler) Detail(c *gin.Context) {

}
