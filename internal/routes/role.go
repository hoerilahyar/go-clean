package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/handler/http"
)

func RegisterRoleRoutes(r *gin.RouterGroup, h *http.RoleHandler) {
	roles := r.Group("/roles")
	{
		roles.GET("", h.GetAll)
		roles.GET("/:id", h.GetByID)
	}
}
