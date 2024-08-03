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

	userHandler := handler.NewUserHandler(u)
	ur := newUserRouter(userHandler)

	articleHandler := handler.NewArticleHandler(a)
	ar := newArticleRouter(articleHandler)

	categoryHandler := handler.NewCategoryHandler(c)
	cr := newCategoryRouter(categoryHandler)

	starHandler := handler.NewStarHandler(s)
	sr := newStarRouter(starHandler)

	advertHandler := handler.NewAdvertHandler(ad)
	adr := newAdvertRouter(advertHandler)

	menuHandler := handler.NewMenuHandler(m)
	mr := newMenuRouter(menuHandler)

	routers = append(routers, ur, ar, cr, sr, adr, mr)
	return
}
