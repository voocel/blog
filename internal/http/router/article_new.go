package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type articleRouterNew struct {
	articleHandler *handler.ArticleHandlerNew
}

func newArticleRouterNew(articleHandler *handler.ArticleHandlerNew) Router {
	return &articleRouterNew{
		articleHandler: articleHandler,
	}
}

func (r *articleRouterNew) Load(engine *gin.Engine) {
	publicGroup := engine.Group("/api/articles")
	{
		publicGroup.GET("", r.articleHandler.GetArticles)
		publicGroup.GET("/:id", r.articleHandler.GetArticle)
	}

	authGroup := engine.Group("/api/articles")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.POST("", r.articleHandler.CreateArticle)
		authGroup.PUT("/:id", r.articleHandler.UpdateArticle)
		authGroup.DELETE("/:id", r.articleHandler.DeleteArticle)
	}
}
