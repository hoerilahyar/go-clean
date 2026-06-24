package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/handler/http"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *http.UserHandler) {
	users := r.Group("/users")
	{
		users.GET("", h.GetAll)
		users.POST("", h.Create)
		users.PUT("", h.Update)
		users.DELETE("/:id", h.Delete)
	}
}
