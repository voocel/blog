package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type articleRouter struct {
	h           *handler.ArticleHandler
	userUseCase *usecase.UserUseCase
}

func newArticleRouter(h *handler.ArticleHandler, userUseCase *usecase.UserUseCase) *articleRouter {
	return &articleRouter{h: h, userUseCase: userUseCase}
}

func (r *articleRouter) Load(g *gin.Engine) {
	group := g.Group("/api/article")
	{
		group.POST("/create", middleware.JWTMiddleware(r.userUseCase), r.h.Create)
		group.GET("/list", r.h.List)
		group.GET("/detail/:aid", r.h.Detail)
		group.PUT("/update", r.h.Update)
		group.PUT("/delete", middleware.JWTMiddleware(r.userUseCase), r.h.DeleteArticlesBatch)
		group.POST("/like/:aid", r.h.Like)
		group.GET("/calendar", r.h.Calendar)
	}
}
