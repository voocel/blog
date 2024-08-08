package router

import (
	"blog/internal/http/handler"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type advertRouter struct {
	h           *handler.AdvertHandler
	userUseCase *usecase.UserUseCase
}

func newAdvertRouter(h *handler.AdvertHandler, userUseCase *usecase.UserUseCase) *advertRouter {
	return &advertRouter{h: h, userUseCase: userUseCase}
}

func (r *advertRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/advert")
	{
		group.POST("/create", r.h.Create)
		group.GET("/list", r.h.List)
		group.PUT("/detail/:aid", r.h.Detail)
		group.PUT("/delete", r.h.DeleteBatch)
	}
}
