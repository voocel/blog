package router

import (
	"blog/internal/http/handler"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type userRouter struct {
	h           *handler.UserHandler
	userUseCase *usecase.UserUseCase
}

func newUserRouter(h *handler.UserHandler, userUseCase *usecase.UserUseCase) *userRouter {
	return &userRouter{h: h, userUseCase: userUseCase}
}

func (r *userRouter) Load(g *gin.Engine) {
	group := g.Group("/api/user")
	{
		group.POST("/login", r.h.Login)
		group.POST("/register", r.h.Register)
		group.GET("/info", r.h.Info)
		group.POST("/logout", r.h.Logout)
		group.GET("/list", r.h.List)
	}
}
