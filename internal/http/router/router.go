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

	userHandler := handler.NewUserHandler(u)
	ur := newUserRouter(userHandler)

	routers = append(routers, ur)
	return
}
