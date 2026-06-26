package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
)

func (h *AuthenticationHandler) Login(c *gin.Context) {

	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	res, err := h.usecase.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
