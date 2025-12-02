package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"intern/internal/configs"
	"intern/internal/core/domain"
	token "intern/pkg/jwt"

	"github.com/gin-gonic/gin"
)

var (
	ErrTokenExpired     = errors.New("token expired")
	ErrPermissionDenied = errors.New("permission denied")
)

// AuthorizerMiddleware checks if the user has permission to access the endpoint
func AuthorizerMiddleware(cfg configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ResponseError{Message: "authorization header is required", Status: http.StatusUnauthorized})
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		if tokenStr == auth {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ResponseError{Message: "invalid token format", Status: http.StatusUnauthorized})
			return
		}

		claims, err := token.ExtractClaims(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ResponseError{Message: "invalid token", Status: http.StatusUnauthorized})
			return
		}

		// Check expiration
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ResponseError{Message: ErrTokenExpired.Error(), Status: http.StatusUnauthorized})
			return
		}

		userID := claims["id"].(string)
		userType := claims["role"].(string)

		// Check permissions for specific endpoints
		path := c.FullPath()
		method := c.Request.Method

		// Check if the endpoint is protected and requires superuser
		isProtected := false
		if method == "POST" && (path == "/api/role" || path == "/api/sysuser") {
			isProtected = true
		}
		if method == "PUT" && strings.HasPrefix(path, "/api/role/") {
			isProtected = true
		}
		if method == "GET" && path == "/api/roles" {
			isProtected = true
		}

		if isProtected {
			if userType != "sysuser" || userID != cfg.SuperUserID {
				c.AbortWithStatusJSON(http.StatusForbidden, domain.ResponseError{Message: ErrPermissionDenied.Error(), Status: http.StatusForbidden})
				return
			}
		}

		c.Set("user_id", userID)
		c.Set("user_type", userType)
		c.Next()
	}
}
