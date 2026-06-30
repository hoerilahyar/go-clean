package authorize

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/bootstrap"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role"
)

// Register authorization route
func RegisterAuthorizationRoutes(
	r *gin.RouterGroup,
	h *bootstrap.AuthorizeHandler,
	service *bootstrap.Services,
) {

	api := r.Group("/iam")

	role.RegisterRoleRoutes(api, h.Role, service)
	permission.RegisterPermissionRoutes(api, h.Permission)
	assignment.RegisterAssignmentRoutes(api, h.Assignment)
}
