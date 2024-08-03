package router

import (
	"blog/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type nenuRouter struct {
	h *handler.MenuHandler
}

func newMenuRouter(h *handler.MenuHandler) *nenuRouter {
	return &nenuRouter{h: h}
}

func (r *nenuRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/menu")
	{
		group.POST("/add", r.h.AddMenu)
		group.GET("/list", r.h.List)
		group.POST("/detail", r.h.Detail)
		group.POST("/update", r.h.UpdateMenu)
		group.POST("/delete", r.h.DeleteMenuById)
		group.POST("/delete_batch", r.h.DeleteMenuBatch)
	}
}
