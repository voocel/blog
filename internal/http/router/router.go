package router

import (
	"blog/internal/http/handler"
	"blog/internal/usecase"
	"blog/internal/usecase/repo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Router 路由接口
type Router interface {
	Load(r *gin.Engine)
}

// GetNewRouters 获取新的路由器列表
func GetNewRouters(db *gorm.DB) (routers []Router) {
	// 创建Repository层
	userRepo := repo.NewUserRepo(db)
	tagRepo := repo.NewTagRepo(db)
	categoryRepo := repo.NewCategoryRepo(db)
	articleRepo := repo.NewArticleRepo(db)
	commentRepo := repo.NewCommentRepo(db)
	discussionRepo := repo.NewDiscussionRepo(db)
	replyRepo := repo.NewReplyRepo(db)
	friendlinkRepo := repo.NewFriendlinkRepo(db)
	statisticsRepo := repo.NewStatisticsRepo(db)

	// 创建UseCase层
	authUseCase := usecase.NewAuthUseCase(userRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)
	tagUseCase := usecase.NewTagUseCase(tagRepo)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	articleUseCase := usecase.NewArticleUseCase(articleRepo, userRepo, categoryRepo, tagRepo)
	commentUseCase := usecase.NewCommentUseCase(commentRepo)
	discussionUseCase := usecase.NewDiscussionUseCase(discussionRepo, replyRepo, userRepo, tagRepo)
	friendlinkUseCase := usecase.NewFriendlinkUseCase(friendlinkRepo)
	statisticsUseCase := usecase.NewStatisticsUseCase(statisticsRepo)

	// 创建Handler层
	authHandler := handler.NewAuthHandler(authUseCase)
	userAdminHandler := handler.NewUserAdminHandler(userUseCase)
	articleHandler := handler.NewArticleHandlerNew(articleUseCase, tagUseCase)
	categoryHandler := handler.NewCategoryHandlerNew(categoryUseCase)
	tagHandler := handler.NewTagHandlerNew(tagUseCase)
	commentHandler := handler.NewCommentHandlerNew(commentUseCase, userUseCase)
	fileHandler := handler.NewFileHandler()
	discussionHandler := handler.NewDiscussionHandler(discussionUseCase)
	friendlinkHandler := handler.NewFriendlinkHandler(friendlinkUseCase)
	statisticsHandler := handler.NewStatisticsHandler(statisticsUseCase)
	systemHandler := handler.NewSystemHandler()

	// 创建Router层
	authRouter := newAuthRouter(authHandler)
	userAdminRouter := newUserAdminRouter(userAdminHandler)
	articleRouter := newArticleRouterNew(articleHandler)
	categoryRouter := newCategoryRouterNew(categoryHandler)
	tagRouter := newTagRouterNew(tagHandler)
	commentRouter := newCommentRouterNew(commentHandler)
	fileRouter := newFileRouter(fileHandler)
	discussionRouter := newDiscussionRouter(discussionHandler)
	friendlinkRouter := newFriendlinkRouter(friendlinkHandler)
	statisticsRouter := newStatisticsRouter(statisticsHandler)
	systemRouter := newSystemRouter(systemHandler)

	routers = append(routers,
		authRouter,
		userAdminRouter,
		articleRouter,
		categoryRouter,
		tagRouter,
		commentRouter,
		fileRouter,
		discussionRouter,
		friendlinkRouter,
		statisticsRouter,
		systemRouter,
	)

	return
}
