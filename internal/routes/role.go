package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/handler"
)

func RegisterRoleRoutes(r *gin.RouterGroup, h *handler.RoleHandler, rolePermission *handler.RolePermissionHandler) {
	roles := r.Group("/roles")
	{
		roles.GET("", h.GetAll)
		// roles.GET("/:id", h.GetByID)
		roles.POST("", h.Create)
		roles.PUT("/:id", h.Update)
		roles.DELETE("/:id", h.Delete)

		roles.GET("/:id/permissions", rolePermission.GetPermissions)
		roles.POST("/:id/permissions", rolePermission.Assign)
		roles.PUT("/:id/permissions", rolePermission.Sync)
		roles.DELETE("/:id/permissions/:permission_id", rolePermission.Remove)

	}
}
