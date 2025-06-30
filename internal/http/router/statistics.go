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

func (r *statisticsRouter) Load(engine *gin.Engine) {
	adminGroup := engine.Group("/api/admin/statistics")
	adminGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminGroup.GET("/dashboard", r.statisticsHandler.GetDashboard)
		adminGroup.GET("/visits", r.statisticsHandler.GetVisits)
	}
} 