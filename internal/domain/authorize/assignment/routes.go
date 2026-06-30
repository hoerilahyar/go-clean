package assignment

import (
	"github.com/gin-gonic/gin"

	assignmentHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/handler"
)

func RegisterAssignmentRoutes(
	r *gin.RouterGroup,
	h *assignmentHandler.AssignmentHandler,
) {

	roles := r.Group("/roles")
	{
		roles.POST("/:id/permissions", h.AssignRolePermissions)
		roles.PUT("/:id/permissions", h.ReplaceRolePermissions)
		roles.GET("/:id/permissions", h.GetRolePermissions)
		roles.DELETE("/:id/permissions/:permission_id", h.DeleteRolePermission)
	}

	users := r.Group("/users")
	{
		users.POST("/:id/roles", h.AssignUserRoles)
		users.PUT("/:id/roles", h.ReplaceUserRoles)
		users.GET("/:id/roles", h.GetUserRoles)
		users.DELETE("/:id/roles/:role_id", h.DeleteUserRole)

		users.POST("/:id/permissions", h.AssignUserPermissions)
		users.PUT("/:id/permissions", h.ReplaceUserPermissions)
		users.GET("/:id/permissions", h.GetUserPermissions)
		users.DELETE("/:id/permissions/:permission_id", h.DeleteUserPermission)
	}

	me := r.Group("/me")
	{
		me.GET("", h.Me)
	}
}
