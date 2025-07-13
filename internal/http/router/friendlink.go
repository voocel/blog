package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type friendlinkRouter struct {
	friendlinkHandler *handler.FriendlinkHandler
}

func newFriendlinkRouter(friendlinkHandler *handler.FriendlinkHandler) Router {
	return &friendlinkRouter{
		friendlinkHandler: friendlinkHandler,
	}
}

func (r *friendlinkRouter) Load(engine *gin.Engine) {
	publicGroup := engine.Group("/api/friendlinks")
	{
		publicGroup.GET("", r.friendlinkHandler.GetFriendlinks)
		publicGroup.GET("/:id", r.friendlinkHandler.GetFriendlink)
	}

	adminGroup := engine.Group("/api/friendlinks")
	adminGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminGroup.POST("", r.friendlinkHandler.CreateFriendlink)
		adminGroup.PUT("/:id", r.friendlinkHandler.UpdateFriendlink)
		adminGroup.DELETE("/:id", r.friendlinkHandler.DeleteFriendlink)
	}
}
