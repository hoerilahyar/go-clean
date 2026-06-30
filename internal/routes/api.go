package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/bootstrap"
	"github.com/hoerilahyar/go-clean/internal/domain/menu"
)

func Register(router *gin.Engine, app *bootstrap.Application) {
	api := router.Group("/api/v1")

	// api.Use(middleware.Auth(app.Config.JWTSecret))

	RegisterUserRoutes(api, app.User, app.Services)

	RegisterAuthorizationRoutes(api, app.Authorize, app.Services)

	RegisterAuthenticationRoutes(api, app.Authentication, app.Services)

	menu.RegisterMenuRoutes(api, app.Menu)
}
