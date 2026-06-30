package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
	response "github.com/hoerilahyar/go-clean/pkg/response"
)

func (h *AuthenticationHandler) Login(c *gin.Context) {

	var req request.LoginRequest

	// Bind request body.
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.BadRequest(err.Error()))
		return
	}

	// Authenticate user.
	res, err := h.usecase.Login(
		c.Request.Context(),
		req,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, res)
}
