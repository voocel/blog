package router

import (
	"blog/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type starRouter struct {
	h *handler.StarHandler
}

func newStarRouter(h *handler.StarHandler) *starRouter {
	return &starRouter{h: h}
}

func (r *starRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/star")
	{
		group.PUT("/add/:aid", r.h.AddStar)
		group.PUT("/remove/:aid", r.h.RemoveStar)
		group.GET("/list", r.h.List)
	}
}
