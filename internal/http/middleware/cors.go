package middleware

import (
	"net/http"
	"strings"

	"blog/config"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		// CORS allowlist:
		// - If http.allowed_origins is configured, only allow listed origins.
		// - If not configured: keep legacy behavior in non-release mode (reflect Origin),
		//   and deny by default in release mode.
		allowed := isOriginAllowed(origin, config.Conf.Http.AllowedOrigins, config.Conf.Mode)
		if allowed && origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT, PATCH")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type, X-Request-ID")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func isOriginAllowed(origin string, allowlist []string, mode string) bool {
	if origin == "" {
		return false
	}
	if len(allowlist) == 0 {
		// Release is stricter by default: if allowlist is empty, deny CORS.
		return strings.ToLower(mode) != "release"
	}
	for _, o := range allowlist {
		if o == origin {
			return true
		}
	}
	return false
}
