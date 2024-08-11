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
	group := g.Group("/api/banner")
	{
		group.POST("/create", middleware.JWTMiddleware(r.userUseCase), r.h.Create)
		group.GET("/detail/:bid", r.h.Detail)
		group.GET("/images", r.h.List)
		group.DELETE("/images", r.h.DeleteBatch)
		group.POST("/images", r.h.CreateBanner)
		group.PUT("/images", middleware.JWTMiddleware(r.userUseCase), r.h.UpdateBanner)
	}

}
