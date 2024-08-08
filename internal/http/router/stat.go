package router

import (
	"blog/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type statRouter struct {
	h *handler.StatHandler
}

func newStatRouter(h *handler.StatHandler) *statRouter {
	return &statRouter{h: h}
}

func (r *statRouter) Load(g *gin.Engine) {
	g.GET("/v1/stat/visit", r.h.Visit)
}
