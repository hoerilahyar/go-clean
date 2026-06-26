package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/bootstrap"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/handler"
)

func RegisterRoleRoutes(
	r *gin.RouterGroup,
	h *handler.RoleHandler,
	service *bootstrap.Services,
) {
	roles := r.Group("/roles")
	{
		roles.GET("", h.GetAll)
		// roles.GET("/:id", h.GetByID)
		roles.POST("", h.Create)
		roles.PUT("/:id", h.Update)
		roles.DELETE("/:id", h.Delete)

	}
}
