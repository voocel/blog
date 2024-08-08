package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type nenuRouter struct {
	h           *handler.MenuHandler
	userUseCase *usecase.UserUseCase
}

func newMenuRouter(h *handler.MenuHandler, userUseCase *usecase.UserUseCase) *nenuRouter {
	return &nenuRouter{h: h, userUseCase: userUseCase}
}

func (r *nenuRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/menu")
	{
		group.POST("/add", middleware.JWTMiddleware(r.userUseCase), r.h.AddMenu)
		group.GET("/list", r.h.List)
		group.GET("/detail/:path", r.h.Detail)
		group.PUT("/update", r.h.UpdateMenu)
		group.POST("/delete/:mid", r.h.DeleteMenuById)
		group.POST("/delete_batch", r.h.DeleteMenuBatch)
	}
}
