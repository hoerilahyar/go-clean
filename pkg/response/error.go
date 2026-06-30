package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Error(c *gin.Context, err error) {

	if err == nil {
		return
	}

	var appErr *apperror.AppError

	if errors.As(err, &appErr) {

		c.JSON(appErr.Status, Response{
			Success: false,
			Error: &ErrorResponse{
				Code:    appErr.Code,
				Message: appErr.Message,
			},
		})

		return
	}

	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Error: &ErrorResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Internal server error",
		},
	})
}
