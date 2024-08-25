package handler

import (
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type StatHandler struct {
	userUsecase *usecase.UserUseCase
}

func NewStatHandler(userUsecase *usecase.UserUseCase) *StatHandler {
	return &StatHandler{userUsecase: userUsecase}
}

func (h *StatHandler) VisitSum(c *gin.Context) {
	resp := new(ApiResponse)
	users, err := h.userUsecase.UserList(c)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	result := map[string]interface{}{
		"article_count":    10,
		"chat_group_count": 20,
		"message_count":    0,
		"now_login_count":  0,
		"now_sign_count":   0,
		"user_count":       len(users),
	}
	resp.Data = result
	c.JSON(http.StatusOK, resp)
	return
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
