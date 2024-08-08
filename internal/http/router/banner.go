package router

import (
	"blog/internal/http/handler"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type bannerRouter struct {
	h           *handler.BannerHandler
	userUseCase *usecase.UserUseCase
}

func newBannerRouter(h *handler.BannerHandler, userUseCase *usecase.UserUseCase) *bannerRouter {
	return &bannerRouter{h: h, userUseCase: userUseCase}
}

func (r *bannerRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/banner")
	{
		group.POST("/create", r.h.Create)
		group.GET("/detail/:bid", r.h.Detail)
		group.GET("/list", r.h.List)
		group.PUT("/update", r.h.Update)
		group.PUT("/delete/:bid", r.h.DeleteById)
		group.PUT("/delete_batch", r.h.DeleteBatch)
	}

}
