package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(
	c *gin.Context,
	data interface{},
	message ...string,
) {

	msg := "Success"

	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

func Created(
	c *gin.Context,
	data interface{},
	message ...string,
) {

	msg := "Created"

	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
