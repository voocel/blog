package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type commentRouterNew struct {
	commentHandler *handler.CommentHandlerNew
}

func newCommentRouterNew(commentHandler *handler.CommentHandlerNew) Router {
	return &commentRouterNew{
		commentHandler: commentHandler,
	}
}

func (r *commentRouterNew) Load(engine *gin.Engine) {
	publicGroup := engine.Group("/api/comments")
	{
		publicGroup.GET("", r.commentHandler.GetComments)
	}

	authGroup := engine.Group("/api/comments")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.POST("", r.commentHandler.CreateComment)
		authGroup.PUT("/:id", r.commentHandler.UpdateComment)
		authGroup.DELETE("/:id", r.commentHandler.DeleteComment)
	}
} 