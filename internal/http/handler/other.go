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
