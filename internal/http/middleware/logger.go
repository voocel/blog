package middleware

import (
	"blog/pkg/log"
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger request logging middleware
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		var bodyBytes []byte
		if c.Request.Body != nil {
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
