package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
	statisticsUseCase *usecase.StatisticsUseCase
}

func NewStatisticsHandler(statisticsUseCase *usecase.StatisticsUseCase) *StatisticsHandler {
	return &StatisticsHandler{
		statisticsUseCase: statisticsUseCase,
	}
}

// GetDashboard 获取仪表盘统计
func (h *StatisticsHandler) GetDashboard(c *gin.Context) {
	statistics, err := h.statisticsUseCase.GetDashboardStatistics(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, "获取统计数据失败"))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(statistics, "获取成功"))
}
