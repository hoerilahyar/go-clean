package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
	response "github.com/hoerilahyar/go-clean/pkg/response"
)

func (h *AuthenticationHandler) ChangePassword(c *gin.Context) {

	var req request.ChangePasswordRequest

	// Bind request body.
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.BadRequest(err.Error()))
		return
	}

	// Retrieve authenticated user ID.
	userID, err := strconv.ParseUint(
		c.GetString("user_id"),
		10,
		64,
	)
	if err != nil {
		response.Error(c, apperror.Unauthorized("Invalid user"))
		return
	}

	// Change user password.
	if err := h.usecase.ChangePassword(
		c.Request.Context(),
		userID,
		req,
	); err != nil {

		response.Error(c, err)
		return
	}

	response.Success(c, nil, "Password changed successfully")
}
