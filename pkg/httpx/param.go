package httpx

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

// ParamString returns a required path parameter.
func ParamString(
	c *gin.Context,
	key string,
) (string, error) {

	value := c.Param(key)

	if value == "" {
		return "", apperror.BadRequest("Missing path parameter: " + key)
	}

	return value, nil
}

// ParamUint64 returns a uint64 path parameter.
func ParamUint64(
	c *gin.Context,
	key string,
) (uint64, error) {

	value, err := ParamString(
		c,
		key,
	)
	if err != nil {
		return 0, err
	}

	id, err := strconv.ParseUint(
		value,
		10,
		64,
	)
	if err != nil {
		return 0, apperror.BadRequest("Invalid path parameter: " + key)
	}

	return id, nil
}

// ParamInt64 returns an int64 path parameter.
func ParamInt64(
	c *gin.Context,
	key string,
) (int64, error) {

	value, err := ParamString(
		c,
		key,
	)
	if err != nil {
		return 0, err
	}

	id, err := strconv.ParseInt(
		value,
		10,
		64,
	)
	if err != nil {
		return 0, apperror.BadRequest("Invalid path parameter: " + key)
	}

	return id, nil
}
