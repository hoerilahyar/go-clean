package bootstrap

import (
	"database/sql"

	"github.com/hoerilahyar/go-clean/internal/config"

	authNHandler "github.com/hoerilahyar/go-clean/internal/domain/authentication/handler"
	assignmentHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/handler"
	permissionHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/handler"
	roleHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/handler"
	menuHandler "github.com/hoerilahyar/go-clean/internal/domain/menu/handler"
	userHandler "github.com/hoerilahyar/go-clean/internal/domain/user/handler"
)

type Application struct {
	Config *config.Config
	DB     *sql.DB

	Services *Services

	User *userHandler.UserHandler

	Authorize *AuthorizeHandler

	Authentication *authNHandler.AuthenticationHandler

	Menu *menuHandler.MenuHandler
}

type AuthorizeHandler struct {
	Role       *roleHandler.RoleHandler
	Permission *permissionHandler.PermissionHandler

	Assignment *assignmentHandler.AssignmentHandler

	Menu *menuHandler.MenuHandler
}
