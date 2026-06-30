package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/domain/menu/handler"
)

func RegisterMenuRoutes(
	r *gin.RouterGroup,
	h *handler.MenuHandler,
) {

	menus := r.Group("/menus")
	{
		menus.GET("", h.GetAll)
		menus.POST("/:id", h.GetByID)
		menus.POST("", h.Create)
		menus.PUT("/:id", h.Update)
		menus.DELETE("/:id", h.Delete)
	}
}
