package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"
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
		group.PUT("/add/:aid", middleware.JWTMiddleware(r.userUseCase), r.h.AddStar)
		group.PUT("/remove/:aid", middleware.JWTMiddleware(r.userUseCase), r.h.RemoveStar)
		group.GET("/list", r.h.List)
	}
}
