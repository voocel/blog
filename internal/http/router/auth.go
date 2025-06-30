package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type authRouter struct {
	authHandler *handler.AuthHandler
}

func newAuthRouter(authHandler *handler.AuthHandler) Router {
	return &authRouter{
		authHandler: authHandler,
	}
}

func (r *authRouter) Load(engine *gin.Engine) {
	authGroup := engine.Group("/api/auth")
	{
		authGroup.POST("/login", r.authHandler.Login)
		authGroup.POST("/register", r.authHandler.Register)
		authGroup.POST("/refresh", r.authHandler.RefreshToken)

		protectedGroup := authGroup.Group("")
		protectedGroup.Use(middleware.AuthMiddleware())
		{
			protectedGroup.POST("/logout", r.authHandler.Logout)
			protectedGroup.GET("/me", r.authHandler.GetProfile)
			protectedGroup.PUT("/profile", r.authHandler.UpdateProfile)
			protectedGroup.POST("/change-password", r.authHandler.ChangePassword)
		}
	}
}
