package router

import (
	"blog/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type advertRouter struct {
	h *handler.AdvertHandler
}

func newAdvertRouter(h *handler.AdvertHandler) *advertRouter {
	return &advertRouter{h: h}
}

func (r *advertRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/article")
	{
		group.POST("/create", r.h.Create)
		group.POST("/list", r.h.List)
		group.GET("/detail/:aid", r.h.Detail)
		group.POST("/delete", r.h.DeleteArticlesBatch)
	}
}
