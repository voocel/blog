package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OtherHandler struct{}

func NewOtherHandler() *OtherHandler {
	return &OtherHandler{}
}

func (h *OtherHandler) News(c *gin.Context) {
	c.JSON(http.StatusOK, ApiResponse{
		Code:    0,
		Message: "ok",
	})
}

func (h *OtherHandler) GetSiteSetting(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"total_view_counts":    10,
			"total_article_counts": 20,
			"site_created_time":    "2018-12-24",
		},
	})
}

func (h *OtherHandler) UpdateSiteSetting(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
	})
}
