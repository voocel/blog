package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type discussionRouter struct {
	discussionHandler *handler.DiscussionHandler
}

func newDiscussionRouter(discussionHandler *handler.DiscussionHandler) Router {
	return &discussionRouter{
		discussionHandler: discussionHandler,
	}
}

func (r *discussionRouter) Load(engine *gin.Engine) {
	publicGroup := engine.Group("/api/discussions")
	{
		publicGroup.GET("", r.discussionHandler.GetDiscussions)
		publicGroup.GET("/:id", r.discussionHandler.GetDiscussion)
	}

	authGroup := engine.Group("/api/discussions")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.POST("", r.discussionHandler.CreateDiscussion)
		authGroup.PUT("/:id", r.discussionHandler.UpdateDiscussion)
		authGroup.DELETE("/:id", r.discussionHandler.DeleteDiscussion)
		authGroup.POST("/:id/replies", r.discussionHandler.CreateReply)
	}
}
