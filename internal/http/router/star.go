package router

import (
	"blog/internal/http/handler"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type starRouter struct {
	h           *handler.StarHandler
	userUseCase *usecase.UserUseCase
}

func newStarRouter(h *handler.StarHandler, userUseCase *usecase.UserUseCase) *starRouter {
	return &starRouter{h: h, userUseCase: userUseCase}
}

func (r *starRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/star")
	{
		group.PUT("/add/:aid", r.h.AddStar)
		group.PUT("/remove/:aid", r.h.RemoveStar)
		group.GET("/list", r.h.List)
	}
}
