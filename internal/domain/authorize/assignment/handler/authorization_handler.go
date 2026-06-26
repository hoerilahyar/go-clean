package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/pkg/utils"
)

func (h *AssignmentHandler) Me(c *gin.Context) {

	userID := c.MustGet("user_id").(uint64)

	result, err := h.usecase.GetMe(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.OK(c, "success", result)
}
