package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type userAdminRouter struct {
	userAdminHandler *handler.UserAdminHandler
}

func newUserAdminRouter(userAdminHandler *handler.UserAdminHandler) Router {
	return &userAdminRouter{
		userAdminHandler: userAdminHandler,
	}
}

func (r *userAdminRouter) Load(engine *gin.Engine) {
	adminGroup := engine.Group("/api/admin/users")
	adminGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminGroup.GET("", r.userAdminHandler.GetUsers)
		adminGroup.POST("", r.userAdminHandler.CreateUser)
		adminGroup.PUT("/:id", r.userAdminHandler.UpdateUser)
		adminGroup.DELETE("/:id", r.userAdminHandler.DeleteUser)
	}
} 