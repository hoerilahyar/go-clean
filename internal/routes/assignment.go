package routes

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
		roles.POST("/:id/permissions", h.AssignPermissionByRoleID)
		roles.PUT("/:id/permissions", h.ReplacePermissionByRoleID)
		roles.GET("/:id/permissions", h.GetPermissionByRoleID)
		roles.DELETE("/:id/permissions/:permission_id", h.RevokePermissionFromRole)
	}

	users := r.Group("/users")
	{
		users.POST("/:id/roles", h.AssignRoleByUserID)
		users.PUT("/:id/roles", h.ReplaceRoleByUserID)
		users.GET("/:id/roles", h.GetRoleByUserID)
		users.DELETE("/:id/roles/:role_id", h.RevokeRoleFromUser)

		users.POST("/:id/permissions", h.AssignPermissionByUserID)
		users.PUT("/:id/permissions", h.ReplacePermissionByUserID)
		users.GET("/:id/permissions", h.GetPermissionByUserID)
		users.DELETE("/:id/permissions/:permission_id", h.RevokePermissionFromUser)
	}

	me := r.Group("/me")
	{
		me.GET("", h.Me)
	}
}
