package router

import (
	"blog/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type userRouter struct {
	h *handler.UserHandler
}

func newUserRouter(h *handler.UserHandler) *userRouter {
	return &userRouter{h: h}
}

func (r *userRouter) Load(g *gin.Engine) {
	ur := g.Group("/v1/user")
	{
		ur.POST("/login", r.h.Login)
		ur.POST("/register", r.h.Register)
		ur.GET("/info", r.h.Info)
	}
}
