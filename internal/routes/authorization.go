package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/bootstrap"
)

// Register authorization route
func RegisterAuthorizationRoutes(
	r *gin.RouterGroup,
	h *bootstrap.AuthorizeHandler,
	service *bootstrap.Services,
) {

	api := r.Group("/iam")

	RegisterRoleRoutes(api, h.Role, service)
	RegisterPermissionRoutes(api, h.Permission)
	RegisterAssignmentRoutes(api, h.Assignment)
}
