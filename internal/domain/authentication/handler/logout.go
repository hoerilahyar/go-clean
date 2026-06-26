package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
)

func (h *AuthenticationHandler) Logout(c *gin.Context) {

	var req request.LogoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.usecase.Logout(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "logout success",
	})
}
