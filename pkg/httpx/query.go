package httpx

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

// QueryString returns a query parameter or a default value.
func QueryString(
	c *gin.Context,
	key string,
	defaultValue string,
) string {

	value := c.Query(key)

	if value == "" {
		return defaultValue
	}

	return value
}

// QueryInt returns an int query parameter.
func QueryInt(
	c *gin.Context,
	key string,
	defaultValue int,
) (int, error) {

	value := c.Query(key)

	if value == "" {
		return defaultValue, nil
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, apperror.BadRequest("Invalid query parameter: " + key)
	}

	return number, nil
}

// QueryInt64 returns an int64 query parameter.
func QueryInt64(
	c *gin.Context,
	key string,
	defaultValue int64,
) (int64, error) {

	value := c.Query(key)

	if value == "" {
		return defaultValue, nil
	}

	number, err := strconv.ParseInt(
		value,
		10,
		64,
	)
	if err != nil {
		return 0, apperror.BadRequest("Invalid query parameter: " + key)
	}

	return number, nil
}

// QueryUint64 returns a uint64 query parameter.
func QueryUint64(
	c *gin.Context,
	key string,
	defaultValue uint64,
) (uint64, error) {

	value := c.Query(key)

	if value == "" {
		return defaultValue, nil
	}

	number, err := strconv.ParseUint(
		value,
		10,
		64,
	)
	if err != nil {
		return 0, apperror.BadRequest("Invalid query parameter: " + key)
	}

	return number, nil
}

// QueryBool returns a boolean query parameter.
func QueryBool(
	c *gin.Context,
	key string,
	defaultValue bool,
) (bool, error) {

	value := c.Query(key)

	if value == "" {
		return defaultValue, nil
	}

	flag, err := strconv.ParseBool(value)
	if err != nil {
		return false, apperror.BadRequest("Invalid query parameter: " + key)
	}

	return flag, nil
}
