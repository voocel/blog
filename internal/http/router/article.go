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
	ur := g.Group("/v1/user")
	{
		ur.POST("/create", r.h.Create)
		ur.POST("/list", r.h.List)
		ur.GET("/detail/:aid", r.h.Detail)
	}
}
