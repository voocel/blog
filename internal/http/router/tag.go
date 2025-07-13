package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type tagRouterNew struct {
	tagHandler *handler.TagHandlerNew
}

func newTagRouterNew(tagHandler *handler.TagHandlerNew) Router {
	return &tagRouterNew{
		tagHandler: tagHandler,
	}
}

func (r *tagRouterNew) Load(engine *gin.Engine) {
	publicGroup := engine.Group("/api/tags")
	{
		publicGroup.GET("", r.tagHandler.GetTags)
		publicGroup.GET("/:id", r.tagHandler.GetTag)
	}

	adminGroup := engine.Group("/api/tags")
	adminGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminGroup.POST("", r.tagHandler.CreateTag)
		adminGroup.PUT("/:id", r.tagHandler.UpdateTag)
		adminGroup.DELETE("/:id", r.tagHandler.DeleteTag)
	}
}
