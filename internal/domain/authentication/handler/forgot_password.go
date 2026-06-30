package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
	response "github.com/hoerilahyar/go-clean/pkg/response"
)

func (h *AuthenticationHandler) ForgotPassword(c *gin.Context) {

	var req request.ForgotPasswordRequest

	// Bind request body.
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.BadRequest(err.Error()))
		return
	}

	// Process forgot password request.
	if err := h.usecase.ForgotPassword(
		c.Request.Context(),
		req,
	); err != nil {

		response.Error(c, err)
		return
	}

	response.Success(c, nil, "Password reset email has been sent")
}
