package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	analyticsUseCase *usecase.AnalyticsUseCase
}

func NewAnalyticsHandler(analyticsUseCase *usecase.AnalyticsUseCase) *AnalyticsHandler {
	return &AnalyticsHandler{analyticsUseCase: analyticsUseCase}
}

// LogVisit - POST /analytics/visit
func (h *AnalyticsHandler) LogVisit(c *gin.Context) {
	var req entity.LogVisitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	if err := h.analyticsUseCase.LogVisit(c.Request.Context(), req, ip, userAgent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// GetLogs - GET /analytics/logs
func (h *AnalyticsHandler) GetLogs(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))

	logs, err := h.analyticsUseCase.GetLogs(c.Request.Context(), startDate, endDate, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

// GetDashboardOverview - GET /analytics/dashboard-overview
func (h *AnalyticsHandler) GetDashboardOverview(c *gin.Context) {
	overview, err := h.analyticsUseCase.GetDashboardOverview(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, overview)
}
