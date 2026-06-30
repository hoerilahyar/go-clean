package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/pkg/httpx"
	response "github.com/hoerilahyar/go-clean/pkg/response"
)

func (h *AssignmentHandler) Me(c *gin.Context) {

	// Retrieve authenticated user ID.
	userID, err := httpx.UserID(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	// Retrieve authenticated user information.
	res, err := h.usecase.GetMe(
		c.Request.Context(),
		userID,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, res)
}
