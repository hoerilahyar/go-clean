package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
)

func (h *AuthenticationHandler) RefreshToken(c *gin.Context) {

	var req request.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	res, err := h.usecase.RefreshToken(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
