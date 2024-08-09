package handler

import (
	"github.com/gin-gonic/gin"
	"time"
)

type StatHandler struct{}

func NewStatHandler() *StatHandler {
	return &StatHandler{}
}

func (h *StatHandler) VisitSum(c *gin.Context) {
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

func (h *StatHandler) VisitWeekLogin(c *gin.Context) {
	var loginDateCountMap = map[string]int{}
	var signDateCountMap = map[string]int{}
	var loginCountList, signCountList []int
	now := time.Now()
	var dateList []string
	// 7 天内
	for i := -6; i <= 0; i++ {
		day := now.AddDate(0, 0, i).Format("2006-01-02")
		dateList = append(dateList, day)
		loginCountList = append(loginCountList, loginDateCountMap[day])
		signCountList = append(signCountList, signDateCountMap[day])
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"date_list":  dateList,
			"login_data": loginCountList,
			"sign_data":  signCountList,
		},
	})
}
