package router

import (
	"blog/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type otherRouter struct {
	h *handler.OtherHandler
}

func newOtherRouter(h *handler.OtherHandler) *otherRouter {
	return &otherRouter{h: h}
}

func (r *otherRouter) Load(g *gin.Engine) {
	group := g.Group("/api/other")
	{
		group.POST("/news", r.h.News)
		group.GET("/get_site_setting", r.h.GetSiteSetting)
		group.POST("/update_site_setting", r.h.UpdateSiteSetting)
	}
}
