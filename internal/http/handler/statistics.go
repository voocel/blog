package handler

import (
	"blog/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
}

func NewStatisticsHandler() *StatisticsHandler {
	return &StatisticsHandler{}
}

// GetDashboard 获取仪表盘统计
func (h *StatisticsHandler) GetDashboard(c *gin.Context) {
	// todo
	dashboard := entity.DashboardStatistics{
		Users: entity.StatisticsItem{
			Total:  1250,
			Growth: 12.5,
		},
		Articles: entity.StatisticsItem{
			Total:  850,
			Growth: 8.3,
		},
		Comments: entity.StatisticsItem{
			Total:  3420,
			Growth: 15.7,
		},
		Visits: entity.StatisticsItem{
			Total:  15680,
			Growth: 22.1,
		},
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(dashboard, "获取成功"))
}

// GetVisits 获取访问统计
func (h *StatisticsHandler) GetVisits(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	// startDate := c.Query("startDate")
	// endDate := c.Query("endDate")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// todo
	var visits []entity.VisitStatistics
	visit := entity.VisitStatistics{
		ID:           1,
		ArticleID:    123,
		ArticleTitle: "示例文章",
		IP:           "192.168.1.1",
		UserAgent:    "Mozilla/5.0 ...",
		Referer:      "https://example.com",
		VisitCount:   1,
		CreatedAt:    "2024-01-01T00:00:00Z",
	}
	visits = append(visits, visit)

	paginatedData := entity.NewPaginatedResponse(visits, 1, page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}
