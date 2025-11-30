package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"
	"blog/internal/usecase"
	"blog/internal/usecase/repo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes 配置所有路由
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// 创建 Repository 层
	userRepo := repo.NewUserRepo(db)
	postRepo := repo.NewPostRepo(db)
	categoryRepo := repo.NewCategoryRepo(db)
	tagRepo := repo.NewTagRepo(db)
	mediaRepo := repo.NewMediaRepo(db)
	analyticsRepo := repo.NewAnalyticsRepo(db)

	// 创建 UseCase 层
	authUseCase := usecase.NewAuthUseCase(userRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)
	postUseCase := usecase.NewPostUseCase(postRepo, categoryRepo, tagRepo)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	tagUseCase := usecase.NewTagUseCase(tagRepo)
	mediaUseCase := usecase.NewMediaUseCase(mediaRepo)
	analyticsUseCase := usecase.NewAnalyticsUseCase(analyticsRepo)

	// 创建 Handler 层
	authHandler := handler.NewAuthHandler(authUseCase, userUseCase)
	userHandler := handler.NewUserHandler(userUseCase)
	postHandler := handler.NewPostHandler(postUseCase)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)
	tagHandler := handler.NewTagHandler(tagUseCase)
	mediaHandler := handler.NewMediaHandler(mediaUseCase)
	analyticsHandler := handler.NewAnalyticsHandler(analyticsUseCase)

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 0. System - Health Check
		v1.GET("/health", handler.HealthCheck)

		// 1. Authentication & User
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
			auth.GET("/me", middleware.JWTAuth(), authHandler.GetCurrentUser)
		}

		users := v1.Group("/users")
		users.Use(middleware.JWTAuth())
		{
			users.PUT("/profile", userHandler.UpdateProfile)
		}

		// 2. Blog Posts
		posts := v1.Group("/posts")
		{
			posts.GET("", postHandler.ListPosts)           // 公开
			posts.GET("/:id", postHandler.GetPost)         // 公开
			posts.POST("", middleware.JWTAuth(), postHandler.CreatePost)
			posts.PUT("/:id", middleware.JWTAuth(), postHandler.UpdatePost)
			posts.DELETE("/:id", middleware.JWTAuth(), postHandler.DeletePost)
		}

		// 3. Taxonomy (Categories & Tags)
		categories := v1.Group("/categories")
		{
			categories.GET("", categoryHandler.ListCategories) // 公开
			categories.POST("", middleware.JWTAuth(), categoryHandler.CreateCategory)
			categories.DELETE("/:id", middleware.JWTAuth(), categoryHandler.DeleteCategory)
		}

		tags := v1.Group("/tags")
		{
			tags.GET("", tagHandler.ListTags) // 公开
			tags.POST("", middleware.JWTAuth(), tagHandler.CreateTag)
			tags.DELETE("/:id", middleware.JWTAuth(), tagHandler.DeleteTag)
		}

		// 4. Media Assets
		v1.POST("/upload", middleware.JWTAuth(), mediaHandler.UploadFile)

		files := v1.Group("/files")
		files.Use(middleware.JWTAuth())
		{
			files.GET("", mediaHandler.ListFiles)
			files.DELETE("/:id", mediaHandler.DeleteFile)
		}

		// 5. Analytics
		analytics := v1.Group("/analytics")
		{
			analytics.POST("/visit", analyticsHandler.LogVisit) // 公开
			analytics.GET("/logs", middleware.JWTAuth(), analyticsHandler.GetLogs)
		}
	}
}
