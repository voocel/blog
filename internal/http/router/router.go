package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(r *gin.Engine, c *Container) {
	// SEO routes (root level)
	r.GET("/sitemap.xml", c.SitemapHandler.GenerateSitemap)
	r.GET("/robots.txt", c.SitemapHandler.RobotsTxt)

	// Page routes — backend injects meta tags into index.html for SEO
	r.GET("/", c.SEOHandler.ServeHome)
	r.GET("/posts", c.SEOHandler.ServePosts)
	r.GET("/post/:slug", c.SEOHandler.ServePost)
	r.GET("/about", c.SEOHandler.ServeAbout)
	r.GET("/clock", c.SEOHandler.ServeFallback)
	r.GET("/settings", c.SEOHandler.ServeFallback)

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
		auth.GET("/me", middleware.JWTAuth(c.UserRepo), c.AuthHandler.GetCurrentUser)
	}
}

func setupUserRoutes(v1 *gin.RouterGroup, c *Container) {
	users := v1.Group("/users")
	users.Use(middleware.JWTAuth(c.UserRepo))
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
		posts.GET("/:slug", c.PostHandler.GetPost)
		posts.GET("/:slug/comments", c.CommentHandler.ListComments)
	}

	// Comments - Authenticated create
	authComments := v1.Group("/posts")
	authComments.Use(middleware.JWTAuth(c.UserRepo))
	{
		authComments.POST("/:slug/comments", c.CommentHandler.CreateComment)
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

	// Likes - Public
	v1.GET("/likes", c.LikeHandler.GetLikes)
	v1.POST("/likes", c.LikeHandler.Like)

	// Analytics - Public tracking
	v1.POST("/analytics/visit", c.AnalyticsHandler.LogVisit)
}

func setupAdminRoutes(v1 *gin.RouterGroup, c *Container) {
	admin := v1.Group("/admin")
	admin.Use(middleware.JWTAuth(c.UserRepo), middleware.AdminOnly())
	{
		setupAdminPostRoutes(admin, c)
		setupAdminTaxonomyRoutes(admin, c)
		setupAdminMediaRoutes(admin, c)
		setupAdminAnalyticsRoutes(admin, c)
		setupAdminEventRoutes(admin, c)
		setupAdminUserRoutes(admin, c)
		setupAdminCommentRoutes(admin, c)
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

func setupAdminUserRoutes(admin *gin.RouterGroup, c *Container) {
	admin.GET("/users", c.UserHandler.ListUsersAdmin)
	admin.PATCH("/users/:id/status", c.UserHandler.UpdateUserStatus)
}

func setupAdminCommentRoutes(admin *gin.RouterGroup, c *Container) {
	admin.GET("/comments", c.CommentHandler.ListAllCommentsAdmin)
	admin.DELETE("/comments/:id", c.CommentHandler.DeleteCommentAdmin)
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
