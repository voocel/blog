package router

import (
	"blog/internal/http/handler"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type categoryRouter struct {
	h           *handler.CategoryHandler
	userUseCase *usecase.UserUseCase
}

func newCategoryRouter(h *handler.CategoryHandler, userUseCase *usecase.UserUseCase) *categoryRouter {
	return &categoryRouter{h: h, userUseCase: userUseCase}
}

func (r *categoryRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/category")
	{
		group.GET("/list", r.h.List)
	}
}
