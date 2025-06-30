package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type systemRouter struct {
	systemHandler *handler.SystemHandler
}

func newSystemRouter(systemHandler *handler.SystemHandler) Router {
	return &systemRouter{
		systemHandler: systemHandler,
	}
}

func (r *systemRouter) Load(engine *gin.Engine) {
	adminGroup := engine.Group("/api/admin/system")
	adminGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminGroup.GET("/info", r.systemHandler.GetSystemInfo)
	}
} 