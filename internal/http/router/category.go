package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type categoryRouterNew struct {
	categoryHandler *handler.CategoryHandlerNew
}

func newCategoryRouterNew(categoryHandler *handler.CategoryHandlerNew) Router {
	return &categoryRouterNew{
		categoryHandler: categoryHandler,
	}
}

func (r *categoryRouterNew) Load(engine *gin.Engine) {
	publicGroup := engine.Group("/api/categories")
	{
		publicGroup.GET("", r.categoryHandler.GetCategories)
		publicGroup.GET("/:id", r.categoryHandler.GetCategory)
	}

	adminGroup := engine.Group("/api/categories")
	adminGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminGroup.POST("", r.categoryHandler.CreateCategory)
		adminGroup.PUT("/:id", r.categoryHandler.UpdateCategory)
		adminGroup.DELETE("/:id", r.categoryHandler.DeleteCategory)
	}
}
