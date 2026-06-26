package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
)

func (h *AuthenticationHandler) ChangePassword(c *gin.Context) {

	var req request.ChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID, err := strconv.ParseUint(
		c.GetString("user_id"),
		10,
		64,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid user",
		})
		return
	}

	if err := h.usecase.ChangePassword(
		c.Request.Context(),
		userID,
		req,
	); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "password changed successfully",
	})
}
