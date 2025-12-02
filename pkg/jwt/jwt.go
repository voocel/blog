package jwt

import (
	"blog/config"
	"blog/internal/entity"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

// Claims represents JWT claims structure
type Claims struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

// TokenPair contains both access and refresh tokens
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"` // Access token expiration in seconds
}

// GenerateTokenPair generates both access and refresh tokens
func GenerateTokenPair(user *entity.User) (*TokenPair, error) {
	accessToken, err := generateToken(user, TokenTypeAccess, time.Duration(config.Conf.App.JwtAccessDuration)*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := generateToken(user, TokenTypeRefresh, time.Duration(config.Conf.App.JwtRefreshDuration)*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(config.Conf.App.JwtAccessDuration * 60), // Convert to seconds
	}, nil
}

// GenerateToken generates a JWT token (deprecated, use GenerateTokenPair instead)
// Kept for backward compatibility
func GenerateToken(user *entity.User) (string, error) {
	return generateToken(user, TokenTypeAccess, time.Duration(config.Conf.App.JwtAccessDuration)*time.Minute)
}

// generateToken is the internal token generation function
func generateToken(user *entity.User, tokenType string, duration time.Duration) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:    user.ID,
		Username:  user.Username,
		Role:      user.Role,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Conf.App.JwtSecret))
}

// ParseToken parses and validates a JWT token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Verify signing algorithm to prevent algorithm downgrade attacks
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Conf.App.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// ValidateRefreshToken validates a refresh token and returns claims if valid
func ValidateRefreshToken(tokenString string) (*Claims, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != TokenTypeRefresh {
		return nil, fmt.Errorf("invalid token type: expected refresh, got %s", claims.TokenType)
	}

	return claims, nil
}
