package handler

import (
	"blog/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error     string `json:"error"`
	Details   string `json:"details,omitempty"`
	RequestID string `json:"requestId,omitempty"`
}

func requestID(c *gin.Context) string {
	if v, ok := c.Get("request_id"); ok {
		if s, ok := v.(string); ok && s != "" {
			return s
		}
	}
	if s := c.GetHeader("X-Request-ID"); s != "" {
		return s
	}
	return ""
}

func shouldExposeDetails(status int) bool {
	// 4xx: safe and useful for debugging on the client side.
	if status >= 400 && status < 500 {
		return true
	}
	// 5xx: avoid leaking internals in production.
	return strings.ToLower(config.Conf.Mode) != "release"
}

// JSONError writes a standardized error response and records error into gin context
// so it can be picked up by request/event loggers via c.Errors.
func JSONError(c *gin.Context, status int, msg string, err error) {
	if err != nil {
		_ = c.Error(err)
	}
	resp := ErrorResponse{
		Error:     msg,
		RequestID: requestID(c),
	}
	if err != nil && shouldExposeDetails(status) {
		resp.Details = err.Error()
	}
	if resp.Error == "" {
		resp.Error = http.StatusText(status)
	}
	c.JSON(status, resp)
}
