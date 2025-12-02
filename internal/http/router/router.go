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
			auth.POST("/refresh", authHandler.RefreshToken) // Refresh access token
			auth.GET("/me", middleware.JWTAuth(), authHandler.GetCurrentUser)
		}

		users := v1.Group("/users")
		users.Use(middleware.JWTAuth())
		{
			users.PUT("/profile", userHandler.UpdateProfile)
		}

		// ========== Public APIs (No Authentication) ==========

		// Blog Posts - Public
		posts := v1.Group("/posts")
		{
			posts.GET("", postHandler.ListPublishedPosts)
			posts.GET("/:id", postHandler.GetPost)
		}

		// Taxonomy (Categories & Tags) - Public Read
		categories := v1.Group("/categories")
		{
			categories.GET("", categoryHandler.ListCategories)
		}

		tags := v1.Group("/tags")
		{
			tags.GET("", tagHandler.ListTags)
		}

		// Analytics - Public tracking
		v1.POST("/analytics/visit", analyticsHandler.LogVisit)

		// ========== Admin APIs (Authentication + Admin Role) ==========

		admin := v1.Group("/admin")
		admin.Use(middleware.JWTAuth(), middleware.AdminOnly())
		{
			// Blog Posts Management
			admin.GET("/posts", postHandler.ListAllPosts)
			admin.GET("/posts/:id", postHandler.GetPostAdmin)
			admin.POST("/posts", postHandler.CreatePost)
			admin.PUT("/posts/:id", postHandler.UpdatePost)
			admin.DELETE("/posts/:id", postHandler.DeletePost)

			// Taxonomy Management
			admin.POST("/categories", categoryHandler.CreateCategory)
			admin.DELETE("/categories/:id", categoryHandler.DeleteCategory)
			admin.POST("/tags", tagHandler.CreateTag)
			admin.DELETE("/tags/:id", tagHandler.DeleteTag)

			// Media Management
			admin.POST("/upload", mediaHandler.UploadFile)
			admin.GET("/files", mediaHandler.ListFiles)
			admin.DELETE("/files/:id", mediaHandler.DeleteFile)

			// Analytics Management
			admin.GET("/analytics/logs", analyticsHandler.GetLogs)
			admin.GET("/analytics/dashboard-overview", analyticsHandler.GetDashboardOverview)
		}
	}
}
