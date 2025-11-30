package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"
	"blog/internal/usecase"
	"blog/internal/usecase/repo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures all routes
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Create Repository layer
	userRepo := repo.NewUserRepo(db)
	postRepo := repo.NewPostRepo(db)
	categoryRepo := repo.NewCategoryRepo(db)
	tagRepo := repo.NewTagRepo(db)
	mediaRepo := repo.NewMediaRepo(db)
	analyticsRepo := repo.NewAnalyticsRepo(db)

	// Create UseCase layer
	authUseCase := usecase.NewAuthUseCase(userRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)
	postUseCase := usecase.NewPostUseCase(postRepo, categoryRepo, tagRepo)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	tagUseCase := usecase.NewTagUseCase(tagRepo)
	mediaUseCase := usecase.NewMediaUseCase(mediaRepo)
	analyticsUseCase := usecase.NewAnalyticsUseCase(analyticsRepo, postRepo, categoryRepo, tagRepo, mediaRepo)

	// Create Handler layer
	authHandler := handler.NewAuthHandler(authUseCase, userUseCase)
	userHandler := handler.NewUserHandler(userUseCase)
	postHandler := handler.NewPostHandler(postUseCase)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)
	tagHandler := handler.NewTagHandler(tagUseCase)
	mediaHandler := handler.NewMediaHandler(mediaUseCase)
	analyticsHandler := handler.NewAnalyticsHandler(analyticsUseCase)

	// API v1 route group
	v1 := r.Group("/api/v1")
	{
		// System - Health Check
		v1.GET("/health", handler.HealthCheck)

		// Authentication & User
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

		// Blog Posts
		posts := v1.Group("/posts")
		{
			posts.GET("", postHandler.ListPosts)   // Public
			posts.GET("/:id", postHandler.GetPost) // Public
			posts.POST("", middleware.JWTAuth(), postHandler.CreatePost)
			posts.PUT("/:id", middleware.JWTAuth(), postHandler.UpdatePost)
			posts.DELETE("/:id", middleware.JWTAuth(), postHandler.DeletePost)
		}

		// Taxonomy (Categories & Tags)
		categories := v1.Group("/categories")
		{
			categories.GET("", categoryHandler.ListCategories) // Public
			categories.POST("", middleware.JWTAuth(), categoryHandler.CreateCategory)
			categories.DELETE("/:id", middleware.JWTAuth(), categoryHandler.DeleteCategory)
		}

		tags := v1.Group("/tags")
		{
			tags.GET("", tagHandler.ListTags) // Public
			tags.POST("", middleware.JWTAuth(), tagHandler.CreateTag)
			tags.DELETE("/:id", middleware.JWTAuth(), tagHandler.DeleteTag)
		}

		// Media Assets
		v1.POST("/upload", middleware.JWTAuth(), mediaHandler.UploadFile)

		files := v1.Group("/files")
		files.Use(middleware.JWTAuth())
		{
			files.GET("", mediaHandler.ListFiles)
			files.DELETE("/:id", mediaHandler.DeleteFile)
		}

		// Analytics
		analytics := v1.Group("/analytics")
		{
			analytics.POST("/visit", analyticsHandler.LogVisit) // Public
			analytics.GET("/logs", middleware.JWTAuth(), analyticsHandler.GetLogs)
			analytics.GET("/dashboard-overview", middleware.JWTAuth(), analyticsHandler.GetDashboardOverview)
		}
	}
}
