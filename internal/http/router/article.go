package router

import (
	"blog/internal/http/handler"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type articleRouter struct {
	h           *handler.ArticleHandler
	userUseCase *usecase.UserUseCase
}

func newArticleRouter(h *handler.ArticleHandler, userUseCase *usecase.UserUseCase) *articleRouter {
	return &articleRouter{h: h, userUseCase: userUseCase}
}

func (r *articleRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/article")
	{
		group.POST("/create", r.h.Create)
		group.GET("/list", r.h.List)
		group.PUT("/detail/:aid", r.h.Detail)
		group.PUT("/delete", r.h.DeleteArticlesBatch)
		group.PUT("/like/:aid", r.h.Like)
	}
}
