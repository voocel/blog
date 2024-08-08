package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TagRouter struct {
	h           *handler.TagHandler
	userUseCase *usecase.UserUseCase
}

func newTagRouter(h *handler.TagHandler, userUseCase *usecase.UserUseCase) *TagRouter {
	return &TagRouter{h: h, userUseCase: userUseCase}
}

func (r *TagRouter) Load(g *gin.Engine) {
	group := g.Group("/v1/tag")
	{
		group.POST("/add", middleware.JWTMiddleware(r.userUseCase), r.h.Create)
		group.GET("/list", r.h.List)
		group.GET("/detail/:tid", r.h.Detail)
		group.PUT("/update", middleware.JWTMiddleware(r.userUseCase), r.h.Update)
		group.PUT("/delete/:aid", middleware.JWTMiddleware(r.userUseCase), r.h.Delete)
		group.PUT("/delete_batch", middleware.JWTMiddleware(r.userUseCase), r.h.DeleteBatch)
	}
}
