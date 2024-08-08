package handler

import "github.com/gin-gonic/gin"

type StatHandler struct{}

func NewStatHandler() *StatHandler {
	return &StatHandler{}
}

func (h *StatHandler) Visit(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"article_count":    10,
			"chat_group_count": 20,
			"message_count":    0,
			"now_login_count":  0,
			"now_sign_count":   0,
			"user_count":       2,
		},
	})
}
