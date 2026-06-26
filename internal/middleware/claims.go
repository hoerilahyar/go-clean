package middleware

import "github.com/gin-gonic/gin"

const (
	ContextUserID   = "user_id"
	ContextUsername = "username"
)

func UserID(c *gin.Context) uint64 {
	value, exists := c.Get(ContextUserID)
	if !exists {
		return 0
	}

	userID, ok := value.(uint64)
	if !ok {
		return 0
	}

	return userID
}

func Username(c *gin.Context) string {
	value, exists := c.Get(ContextUsername)
	if !exists {
		return ""
	}

	username, ok := value.(string)
	if !ok {
		return ""
	}

	return username
}
