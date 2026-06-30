package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/go-clean/internal/bootstrap"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize"
	"github.com/hoerilahyar/go-clean/internal/domain/menu"
	"github.com/hoerilahyar/go-clean/internal/domain/user"
)

func Register(router *gin.Engine, app *bootstrap.Application) {
	api := router.Group("/api/v1")

	// api.Use(middleware.Auth(app.Config.JWTSecret))

	user.RegisterUserRoutes(api, app.User, app.Services)

	authorize.RegisterAuthorizationRoutes(api, app.Authorize, app.Services)

	authentication.RegisterAuthenticationRoutes(api, app.Authentication, app.Services)

	menu.RegisterMenuRoutes(api, app.Menu)
}
