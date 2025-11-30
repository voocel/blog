package middleware

import (
	"blog/pkg/log"
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger 请求日志中间件
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// 读取请求体（用于日志记录）
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			// 重新设置请求体，因为读取后会被消耗
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// 处理请求
		c.Next()

		// 计算耗时
		latency := time.Since(startTime)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		// 记录详细日志
		if statusCode >= 400 {
			// 错误请求记录更详细的信息
			log.Errorw("HTTP Request Error",
				log.Pair("status", statusCode),
				log.Pair("method", method),
				log.Pair("path", path),
				log.Pair("ip", clientIP),
				log.Pair("latency", latency.String()),
				log.Pair("user-agent", c.Request.UserAgent()),
				log.Pair("body", string(bodyBytes)),
				log.Pair("errors", c.Errors.String()),
			)
		} else {
			// 正常请求只记录基本信息
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
