package middleware

import (
	"blog/pkg/log"
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const maxLogBodyBytes = 64 * 1024 // 64KB

// RequestLogger request logging middleware
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		var bodyBytes []byte
		if c.Request.Body != nil &&
			!strings.HasPrefix(path, "/api/v1/auth/") &&
			strings.Contains(strings.ToLower(c.GetHeader("Content-Type")), "application/json") &&
			c.Request.ContentLength > 0 && c.Request.ContentLength <= maxLogBodyBytes {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			// Reset request body as it's consumed after reading
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		c.Next()

		latency := time.Since(startTime)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		if statusCode >= 400 {
			log.Errorw("HTTP Request Error",
				log.Pair("status", statusCode),
				log.Pair("method", method),
				log.Pair("path", path),
				log.Pair("ip", clientIP),
				log.Pair("latency", latency.String()),
				log.Pair("body", string(bodyBytes)),
				log.Pair("errors", c.Errors.String()),
			)
		} else {
			// Skip logging for health check endpoint
			if path == "/api/v1/health" {
				return
			}
			log.Infow("HTTP Request",
				log.Pair("status", statusCode),
				log.Pair("method", method),
				log.Pair("path", path),
				log.Pair("ip", clientIP),
				log.Pair("latency", latency.String()),
			)
		}
	}
}
