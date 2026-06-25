package routes

import (
	"github.com/gin-gonic/gin"

	userHandler "github.com/hoerilahyar/go-clean/internal/domain/user/handler"
)

func RegisterUserRoutes(
	r *gin.RouterGroup,
	user *userHandler.UserHandler,
	// userRole *userHandler.UserRoleHandler,
	// userPermission *userHandler.UserPermissionHandler,
) {
	users := r.Group("/users")
	{
		// CRUD User
		users.GET("", user.GetAll)
		// users.GET("/:id", user.GetByID)
		users.POST("", user.Create)
		users.PUT("/:id", user.Update)
		users.DELETE("/:id", user.Delete)

		// User Roles
		// users.GET("/:id/roles", userRole.GetRoles)
		// users.POST("/:id/roles", userRole.Assign)
		// users.PUT("/:id/roles", userRole.Sync)
		// users.DELETE("/:id/roles/:role_id", userRole.Remove)

		// // User Permissions
		// users.GET("/:id/permissions", userPermission.GetPermissions)
		// users.POST("/:id/permissions", userPermission.Assign)
		// users.PUT("/:id/permissions", userPermission.Sync)
		// users.DELETE("/:id/permissions/:permission_id", userPermission.Remove)
	}
}
