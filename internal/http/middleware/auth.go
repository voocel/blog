package middleware

import (
	"blog/internal/entity"
	"blog/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := getTokenFromHeader(c)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "缺少认证令牌"))
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "无效的认证令牌"))
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}

// AdminMiddleware 管理员权限中间件
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, entity.NewErrorResponse(401, "用户信息不存在"))
			c.Abort()
			return
		}

		if role != "admin" {
			c.JSON(http.StatusForbidden, entity.NewErrorResponse(403, "权限不足"))
			c.Abort()
			return
		}

		c.Next()
	}
}

// OptionalAuthMiddleware 可选认证中间件
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := getTokenFromHeader(c)
		if tokenString != "" {
			claims, err := jwt.ValidateToken(tokenString)
			if err == nil {
				c.Set("user_id", claims.UserID)
				c.Set("username", claims.Username)
				c.Set("user_role", claims.Role)
			}
		}
		c.Next()
	}
}

// getTokenFromHeader 从请求头中获取令牌
func getTokenFromHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	// Bearer token格式
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}

	return ""
}

// GetCurrentUserID 获取当前用户ID
func GetCurrentUserID(c *gin.Context) (int64, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	return userID.(int64), true
}

// GetCurrentUserRole 获取当前用户角色
func GetCurrentUserRole(c *gin.Context) (string, bool) {
	role, exists := c.Get("user_role")
	if !exists {
		return "", false
	}
	return role.(string), true
}
