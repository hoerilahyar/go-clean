package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
	response "github.com/hoerilahyar/go-clean/pkg/response"
)

func (h *AuthenticationHandler) Logout(c *gin.Context) {

	var req request.LogoutRequest

	// Bind request body.
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.BadRequest(err.Error()))
		return
	}

	// Logout current session.
	if err := h.usecase.Logout(
		c.Request.Context(),
		req,
	); err != nil {

		response.Error(c, err)
		return
	}

	response.Success(c, nil, "Logout successful")
}
