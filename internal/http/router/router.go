package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(r *gin.Engine, c *Container) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", handler.HealthCheck)

		// Setup route modules
		setupAuthRoutes(v1, c)
		setupUserRoutes(v1, c)
		setupPublicRoutes(v1, c)
		setupAdminRoutes(v1, c)
	}
}

func setupAuthRoutes(v1 *gin.RouterGroup, c *Container) {
	auth := v1.Group("/auth")
	{
		auth.POST("/login", c.AuthHandler.Login)
		auth.POST("/register", c.AuthHandler.Register)
		auth.POST("/refresh", c.AuthHandler.RefreshToken)
		auth.GET("/me", middleware.JWTAuth(), c.AuthHandler.GetCurrentUser)
	}
}

func setupUserRoutes(v1 *gin.RouterGroup, c *Container) {
	users := v1.Group("/users")
	users.Use(middleware.JWTAuth())
	{
		users.PUT("/profile", c.UserHandler.UpdateProfile)
		users.POST("/avatar", c.MediaHandler.UploadAvatar)
	}
}

func setupPublicRoutes(v1 *gin.RouterGroup, c *Container) {
	// Blog Posts - Public
	posts := v1.Group("/posts")
	{
		posts.GET("", c.PostHandler.ListPublishedPosts)
		posts.GET("/:id", c.PostHandler.GetPost)
	}

	// Taxonomy - Public Read
	categories := v1.Group("/categories")
	{
		categories.GET("", c.CategoryHandler.ListCategories)
	}

	tags := v1.Group("/tags")
	{
		tags.GET("", c.TagHandler.ListTags)
	}

	// Analytics - Public tracking
	v1.POST("/analytics/visit", c.AnalyticsHandler.LogVisit)
}

func setupAdminRoutes(v1 *gin.RouterGroup, c *Container) {
	admin := v1.Group("/admin")
	admin.Use(middleware.JWTAuth(), middleware.AdminOnly())
	{
		setupAdminPostRoutes(admin, c)
		setupAdminTaxonomyRoutes(admin, c)
		setupAdminMediaRoutes(admin, c)
		setupAdminAnalyticsRoutes(admin, c)
		setupAdminEventRoutes(admin, c)
	}
}

func setupAdminPostRoutes(admin *gin.RouterGroup, c *Container) {
	admin.GET("/posts", c.PostHandler.ListAllPosts)
	admin.GET("/posts/:id", c.PostHandler.GetPostAdmin)
	admin.POST("/posts", c.PostHandler.CreatePost)
	admin.PUT("/posts/:id", c.PostHandler.UpdatePost)
	admin.DELETE("/posts/:id", c.PostHandler.DeletePost)
}

func setupAdminTaxonomyRoutes(admin *gin.RouterGroup, c *Container) {
	// Categories
	admin.POST("/categories", c.CategoryHandler.CreateCategory)
	admin.DELETE("/categories/:id", c.CategoryHandler.DeleteCategory)

	// Tags
	admin.POST("/tags", c.TagHandler.CreateTag)
	admin.DELETE("/tags/:id", c.TagHandler.DeleteTag)
}

func setupAdminMediaRoutes(admin *gin.RouterGroup, c *Container) {
	admin.POST("/upload", c.MediaHandler.UploadFile)
	admin.GET("/files", c.MediaHandler.ListFiles)
	admin.DELETE("/files/:id", c.MediaHandler.DeleteFile)
}

func setupAdminAnalyticsRoutes(admin *gin.RouterGroup, c *Container) {
	admin.GET("/analytics/logs", c.AnalyticsHandler.GetLogs)
	admin.GET("/analytics/dashboard-overview", c.AnalyticsHandler.GetDashboardOverview)
}

func setupAdminEventRoutes(admin *gin.RouterGroup, c *Container) {
	// General event queries
	admin.GET("/events", c.SystemEventHandler.ListEvents)
	admin.GET("/events/user/:id", c.SystemEventHandler.GetUserEvents)
	admin.GET("/events/trace/:request_id", c.SystemEventHandler.GetRequestTrace)
	admin.GET("/events/type/:event_type", c.SystemEventHandler.GetEventsByType)

	// Convenience endpoints for specific event types
	admin.GET("/events/audit", c.SystemEventHandler.GetAuditLogs)
	admin.GET("/events/security", c.SystemEventHandler.GetSecurityEvents)
	admin.GET("/events/errors", c.SystemEventHandler.GetSystemErrors)
}
