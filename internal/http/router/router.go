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

	userHandler := handler.NewUserHandler(u)
	ur := newUserRouter(userHandler)

	articleHandler := handler.NewArticleHandler(a)
	ar := newArticleRouter(articleHandler)

	routers = append(routers, ur)
	return
}
