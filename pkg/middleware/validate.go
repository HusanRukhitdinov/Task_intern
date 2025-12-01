package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"intern/internal/configs"
	"intern/pkg/logger"
)

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaims(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(configs.SignKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("faild to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func GetIdByClaims(c *gin.Context, log logger.ILogger) (string, error) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		log.Error("user not found")
		return "", fmt.Errorf("user not found")
	}

	claims, ok := claimsRaw.(jwt.MapClaims)
	if !ok {
		log.Error("claims is not a valid map")
		return "", fmt.Errorf("claims is not a valid map")
	}

	userId, ok := claims["id"].(string)
	if !ok {
		log.Error("user_id not found or is not a string")
		return "", fmt.Errorf("user_id not found or is not a string")
	}
	return userId, nil
}
