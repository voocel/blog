package router

import (
	"blog/internal/http/handler"
	"blog/internal/usecase"
	"blog/internal/usecase/repo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router interface {
	Load(r *gin.Engine)
}

func GetRouters(db *gorm.DB) (routers []Router) {
	u := usecase.NewUserUseCase(repo.NewUserRepo(db))
	a := usecase.NewArticleUseCase(repo.NewArticleRepo(db))
	c := usecase.NewCategoryUseCase(repo.NewCategoryRepo(db))
	s := usecase.NewStarUseCase(repo.NewStarRepo(db))
	ad := usecase.NewAdvertUseCase(repo.NewAdvertRepo(db))
	m := usecase.NewMenuUseCase(repo.NewMenuRepo(db))
	b := usecase.NewBannerUseCase(repo.NewBannerRepo(db))
	cm := usecase.NewCommentUseCase(repo.NewCommentRepo(db))

	userHandler := handler.NewUserHandler(u)
	ur := newUserRouter(userHandler, u)

	articleHandler := handler.NewArticleHandler(a)
	ar := newArticleRouter(articleHandler, u)

	categoryHandler := handler.NewCategoryHandler(c)
	cr := newCategoryRouter(categoryHandler, u)

	starHandler := handler.NewStarHandler(s)
	sr := newStarRouter(starHandler, u)

	advertHandler := handler.NewAdvertHandler(ad)
	adr := newAdvertRouter(advertHandler, u)

	menuHandler := handler.NewMenuHandler(m)
	mr := newMenuRouter(menuHandler, u)

	bannerHandler := handler.NewBannerHandler(b)
	br := newBannerRouter(bannerHandler, u)

	commentHandler := handler.NewCommentHandler(cm)
	cmr := newCommentRouter(commentHandler, u)

	routers = append(routers, ur, ar, cr, sr, adr, mr, br, cmr)
	return
}
