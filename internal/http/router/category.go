package router

import (
	"blog/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type categoryRouter struct {
	h *handler.CategoryHandler
}

func newCategoryRouter(h *handler.CategoryHandler) *categoryRouter {
	return &categoryRouter{h: h}
}

func (r *categoryRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/category")
	{
		group.GET("/list", r.h.List)
	}
}
