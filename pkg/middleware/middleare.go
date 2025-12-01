package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"intern/internal/core/domain"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ResponseError{Error: fmt.Errorf("authorization header is required").Error()})
			return
		}
		valid, err := ValidateToken(auth)
		if err != nil || !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ResponseError{Error: "invalid token:"})
			return
		}

		claims, err := ExtractClaims(auth)
		if err != nil || !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ResponseError{Error: "invalid token claims:"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
