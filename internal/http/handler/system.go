package handler

import (
	"blog/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemHandler struct {
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}

// GetSystemInfo 获取系统信息
func (h *SystemHandler) GetSystemInfo(c *gin.Context) {
	// todo
	systemInfo := entity.SystemInfo{
		Language:  "Chinese",
		Version:   "1.0.0",
		WebServer: "Gin/1.9.1",
		Domain:    "blog.example.com",
		IP:        "127.0.0.1",
		UserAgent: c.GetHeader("User-Agent"),
		Database: entity.DatabaseInfo{
			Type:    "PostgreSQL",
			Version: "8.0.33",
		},
		PHP: &entity.PHPInfo{
			Version:    "8.2.0",
			Extensions: []string{"json", "mbstring", "openssl", "PDO", "curl"},
		},
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(systemInfo, "获取成功"))
} 