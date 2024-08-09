package handler

import (
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogstashHandler struct {
	useCase *usecase.LogstashUseCase
}

func NewLogstashHandler(u *usecase.LogstashUseCase) *LogstashHandler {
	return &LogstashHandler{
		useCase: u,
	}
}

func (h *LogstashHandler) List(c *gin.Context) {
	resp := new(ApiResponse)
	result, err := h.useCase.List(c)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = result
	c.JSON(http.StatusOK, resp)
	return
}
