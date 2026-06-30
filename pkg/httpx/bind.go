package httpx

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/pkg/apperror"
	"github.com/hoerilahyar/go-clean/pkg/validator"
)

// BindJSON binds and validates a JSON request body.
func BindJSON[T any](
	c *gin.Context,
) (T, error) {

	var req T

	if err := c.ShouldBindJSON(&req); err != nil {
		var zero T
		return zero, apperror.BadRequest(err.Error())
	}

	if err := validator.Validate(req); err != nil {
		var zero T
		return zero, apperror.BadRequest(err.Error())
	}

	return req, nil
}

// BindQuery binds and validates query parameters.
func BindQuery[T any](
	c *gin.Context,
) (T, error) {

	var req T

	if err := c.ShouldBindQuery(&req); err != nil {
		var zero T
		return zero, apperror.BadRequest(err.Error())
	}

	if err := validator.Validate(req); err != nil {
		var zero T
		return zero, apperror.BadRequest(err.Error())
	}

	return req, nil
}
