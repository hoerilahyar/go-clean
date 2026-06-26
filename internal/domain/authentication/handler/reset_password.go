package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
)

func (h *AuthenticationHandler) ResetPassword(c *gin.Context) {

	var req request.ResetPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.usecase.ResetPassword(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "password reset successfully",
	})
}
