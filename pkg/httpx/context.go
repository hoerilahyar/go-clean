package httpx

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func UserID(
	c *gin.Context,
) (uint64, error) {

	value, exists := c.Get("user_id")
	if !exists {
		return 0, apperror.Unauthorized("Unauthorized")
	}

	userID, ok := value.(uint64)
	if !ok {
		return 0, apperror.Unauthorized("Invalid user ID")
	}

	if userID == 0 {
		return 0, apperror.Unauthorized("Invalid user ID")
	}

	return userID, nil
}
