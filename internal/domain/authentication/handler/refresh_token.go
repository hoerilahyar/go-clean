package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
	response "github.com/hoerilahyar/go-clean/pkg/response"
)

func (h *AuthenticationHandler) RefreshToken(c *gin.Context) {

	var req request.RefreshTokenRequest

	// Bind request body.
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.BadRequest(err.Error()))
		return
	}

	// Refresh access token.
	res, err := h.usecase.RefreshToken(
		c.Request.Context(),
		req,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, res)
}
