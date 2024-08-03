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
	group := g.Group("/v1/user")
	{
		group.POST("/login", r.h.Login)
		group.POST("/register", r.h.Register)
		group.GET("/info", r.h.Info)
	}
}
