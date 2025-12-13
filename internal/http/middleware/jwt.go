package middleware

import (
	"blog/internal/usecase"
	"blog/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth validates JWT access tokens
func JWTAuth(userRepo usecase.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Includes signature verification and algorithm validation
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Ensure it's an access token (not a refresh token)
		if claims.TokenType != jwt.TokenTypeAccess {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token type"})
			c.Abort()
			return
		}

		// Enforce user status and token version (revocation).
		user, err := userRepo.GetByID(c.Request.Context(), claims.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		if user.Status == "banned" {
			c.JSON(http.StatusForbidden, gin.H{"error": "User is banned"})
			c.Abort()
			return
		}

		tokenTV := claims.TokenVersion
		if tokenTV <= 0 {
			tokenTV = 1 // Backward compatibility for older tokens without tv
		}
		userTV := user.TokenVersion
		if userTV <= 0 {
			userTV = 1
		}
		if tokenTV != userTV {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token revoked"})
			c.Abort()
			return
		}

		// Use DB values so role/status changes take effect immediately.
		c.Set("user_id", user.ID)
		c.Set("username", user.Username)
		c.Set("role", user.Role)

		c.Next()
	}
}
