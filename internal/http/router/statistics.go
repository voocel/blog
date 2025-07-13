package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type statisticsRouter struct {
	statisticsHandler *handler.StatisticsHandler
}

func newStatisticsRouter(statisticsHandler *handler.StatisticsHandler) Router {
	return &statisticsRouter{
		statisticsHandler: statisticsHandler,
	}
}

func (r *statisticsRouter) Load(rg *gin.Engine) {
	v1 := rg.Group("/api")
	adminGroup := v1.Group("/admin/statistics")
	adminGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())

	adminGroup.GET("/dashboard", r.statisticsHandler.GetDashboard)
}
