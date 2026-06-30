package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/handler"
)

func RegisterPermissionRoutes(r *gin.RouterGroup, h *handler.PermissionHandler) {
	permissions := r.Group("/permissions")
	{
		permissions.GET("", h.GetAll)
		// permissions.GET("/:id", h.GetByID)
		permissions.GET("/group/:group", h.GetByGroup)

		permissions.POST("", h.Create)
		permissions.PUT("/:id", h.Update)
		permissions.DELETE("/:id", h.Delete)
	}
}
