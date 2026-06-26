package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/pkg/jwt"
)

func Auth(jwtService jwt.Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "authorization header is required",
			})
			return
		}

		const bearer = "Bearer "

		if !strings.HasPrefix(authHeader, bearer) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authorization header",
			})
			return
		}

		token := strings.TrimPrefix(authHeader, bearer)

		claims, err := jwtService.ValidateAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid access token",
			})
			return
		}

		c.Set(ContextUserID, claims.UserID)
		c.Set(ContextUsername, claims.Username)

		c.Next()
	}
}
