package router

import (
	"blog/internal/http/handler"
	"blog/internal/usecase"
	"blog/internal/usecase/repo"

	"gorm.io/gorm"
)

// Container holds all application dependencies
// This pattern provides centralized dependency management and makes testing easier
type Container struct {
	// Repositories
	UserRepo        usecase.UserRepo
	PostRepo        usecase.PostRepo
	CategoryRepo    usecase.CategoryRepo
	TagRepo         usecase.TagRepo
	MediaRepo       usecase.MediaRepo
	AnalyticsRepo   usecase.AnalyticsRepo
	SystemEventRepo usecase.SystemEventRepo
	CommentRepo     usecase.CommentRepo
	LikeRepo        usecase.LikeRepo

	// UseCases
	AuthUseCase        *usecase.AuthUseCase
	UserUseCase        *usecase.UserUseCase
	PostUseCase        *usecase.PostUseCase
	CategoryUseCase    *usecase.CategoryUseCase
	TagUseCase         *usecase.TagUseCase
	MediaUseCase       *usecase.MediaUseCase
	AnalyticsUseCase   *usecase.AnalyticsUseCase
	SystemEventUseCase *usecase.SystemEventUseCase
	CommentUseCase     *usecase.CommentUseCase
	LikeUseCase        *usecase.LikeUseCase

	// Handlers
	AuthHandler        *handler.AuthHandler
	UserHandler        *handler.UserHandler
	PostHandler        *handler.PostHandler
	CategoryHandler    *handler.CategoryHandler
	TagHandler         *handler.TagHandler
	MediaHandler       *handler.MediaHandler
	AnalyticsHandler   *handler.AnalyticsHandler
	SystemEventHandler *handler.SystemEventHandler
	CommentHandler     *handler.CommentHandler
	LikeHandler        *handler.LikeHandler
	SitemapHandler     *handler.SitemapHandler
}

// NewContainer creates and initializes all application dependencies
// Dependencies are created in order: Repository -> UseCase -> Handler
func NewContainer(db *gorm.DB) *Container {
	c := &Container{}

	// Initialize Repositories
	c.UserRepo = repo.NewUserRepo(db)
	c.PostRepo = repo.NewPostRepo(db)
	c.CategoryRepo = repo.NewCategoryRepo(db)
	c.TagRepo = repo.NewTagRepo(db)
	c.MediaRepo = repo.NewMediaRepo(db)
	c.AnalyticsRepo = repo.NewAnalyticsRepo(db)
	c.SystemEventRepo = repo.NewSystemEventRepo(db)
	c.CommentRepo = repo.NewCommentRepo(db)
	c.LikeRepo = repo.NewLikeRepo(db)

	// Initialize UseCases
	c.AuthUseCase = usecase.NewAuthUseCase(c.UserRepo)
	c.UserUseCase = usecase.NewUserUseCase(c.UserRepo)
	c.PostUseCase = usecase.NewPostUseCase(c.PostRepo, c.CategoryRepo, c.TagRepo, c.AnalyticsRepo)
	c.CategoryUseCase = usecase.NewCategoryUseCase(c.CategoryRepo)
	c.TagUseCase = usecase.NewTagUseCase(c.TagRepo)
	c.MediaUseCase = usecase.NewMediaUseCase(c.MediaRepo)
	c.AnalyticsUseCase = usecase.NewAnalyticsUseCase(c.AnalyticsRepo, c.PostRepo, c.CategoryRepo, c.TagRepo, c.MediaRepo)
	c.SystemEventUseCase = usecase.NewSystemEventUseCase(c.SystemEventRepo)
	c.CommentUseCase = usecase.NewCommentUseCase(c.CommentRepo, c.PostRepo, c.UserRepo)
	c.LikeUseCase = usecase.NewLikeUseCase(c.LikeRepo)

	// Initialize Handlers
	c.AuthHandler = handler.NewAuthHandler(c.AuthUseCase, c.UserUseCase)
	c.UserHandler = handler.NewUserHandler(c.UserUseCase)
	c.PostHandler = handler.NewPostHandler(c.PostUseCase)
	c.CategoryHandler = handler.NewCategoryHandler(c.CategoryUseCase)
	c.TagHandler = handler.NewTagHandler(c.TagUseCase)
	c.MediaHandler = handler.NewMediaHandler(c.MediaUseCase)
	c.AnalyticsHandler = handler.NewAnalyticsHandler(c.AnalyticsUseCase)
	c.SystemEventHandler = handler.NewSystemEventHandler(c.SystemEventUseCase)
	c.CommentHandler = handler.NewCommentHandler(c.CommentUseCase, c.PostUseCase)
	c.LikeHandler = handler.NewLikeHandler(c.LikeUseCase)
	c.SitemapHandler = handler.NewSitemapHandler(c.PostRepo)

	return c
}
