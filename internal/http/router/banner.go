package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"
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
		group.POST("/create", middleware.JWTMiddleware(r.userUseCase), r.h.Create)
		group.GET("/detail/:bid", r.h.Detail)
		group.GET("/list", r.h.List)
		group.PUT("/update", middleware.JWTMiddleware(r.userUseCase), r.h.Update)
		group.PUT("/delete/:bid", middleware.JWTMiddleware(r.userUseCase), r.h.DeleteById)
		group.PUT("/delete_batch", middleware.JWTMiddleware(r.userUseCase), r.h.DeleteBatch)
	}

}
