package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type commentRouter struct {
	h           *handler.CommentHandler
	userUseCase *usecase.UserUseCase
}

func newCommentRouter(h *handler.CommentHandler, userUseCase *usecase.UserUseCase) *commentRouter {
	return &commentRouter{h: h, userUseCase: userUseCase}
}

func (r *commentRouter) Load(g *gin.Engine) {
	group := g.Group("/api/comment")
	{
		group.POST("/create", middleware.JWTMiddleware(r.userUseCase), r.h.Create)
		group.GET("/article_list/:aid", r.h.GetArticleCommentList)
		group.GET("/list", r.h.GetAllCommentList)
		group.POST("/delete/:cid", middleware.JWTMiddleware(r.userUseCase), r.h.Delete)
	}
}
