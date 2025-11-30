package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheck - GET /health
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
