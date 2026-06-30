package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
	TotalPage int   `json:"total_page"`
}

func SuccessWithPagination(
	c *gin.Context,
	data interface{},
	meta Pagination,
) {

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Success",
		Data:    data,
		Meta:    meta,
	})
}
