package router

import (
	"blog/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type articleRouter struct {
	h *handler.ArticleHandler
}

func newArticleRouter(h *handler.ArticleHandler) *articleRouter {
	return &articleRouter{h: h}
}

func (r *articleRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/article")
	{
		group.POST("/create", r.h.Create)
		group.POST("/list", r.h.List)
		group.GET("/detail/:aid", r.h.Detail)
		group.POST("/delete", r.h.DeleteArticlesBatch)
		group.PUT("/delete/:aid", r.h.Like)
	}
}
