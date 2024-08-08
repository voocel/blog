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
	group := g.Group("/v1/other")
	{
		group.POST("/news", r.h.News)
	}
}
