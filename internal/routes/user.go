package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/bootstrap"
	userHandler "github.com/hoerilahyar/go-clean/internal/domain/user/handler"
	"github.com/hoerilahyar/go-clean/internal/middleware"
)

func RegisterUserRoutes(
	r *gin.RouterGroup,
	user *userHandler.UserHandler,
	service *bootstrap.Services,
) {

	users := r.Group("/users")
	users.Use(middleware.Auth(service.JWT))
	{
		// CRUD User
		users.GET("", user.GetAll)
		users.POST("", user.Create)
		users.PUT("/:id", user.Update)
		users.DELETE("/:id", user.Delete)
	}
}
