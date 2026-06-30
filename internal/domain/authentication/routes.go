package authentication

import (
	"github.com/gin-gonic/gin"

	"github.com/hoerilahyar/go-clean/internal/bootstrap"
	authHandler "github.com/hoerilahyar/go-clean/internal/domain/authentication/handler"
	"github.com/hoerilahyar/go-clean/internal/middleware"
)

func RegisterAuthenticationRoutes(
	r *gin.RouterGroup,
	auth *authHandler.AuthenticationHandler,
	service *bootstrap.Services,
) {
	authN := r.Group("/auth")

	// Public endpoint
	authN.POST("/login", auth.Login)
	authN.POST("/refresh", auth.RefreshToken)

	password := authN.Group("/password")
	{
		password.POST("/forgot", auth.ForgotPassword)
		password.POST("/reset", auth.ResetPassword)
	}

	// Protected endpoint
	protected := authN.Group("")
	protected.Use(middleware.Auth(service.JWT))
	{
		protected.POST("/logout", auth.Logout)
		protected.POST("/password/change", auth.ChangePassword)
	}
}
