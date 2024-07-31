package router

import (
	"blog/internal/http/handler"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Load(r *gin.Engine)
}

func GetRouters() (routers []Router) {
	u := usecase.NewUserUseCase()

	userHandler := handler.NewUserHandler(u)
	ur := newUserRouter(userHandler)

	routers = append(routers, ur)
	return
}
