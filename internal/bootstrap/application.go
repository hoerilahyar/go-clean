package bootstrap

import (
	"database/sql"

	"github.com/hoerilahyar/go-clean/internal/config"

	assignmentHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/handler"
	permissionHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/handler"
	roleHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/handler"
	userHandler "github.com/hoerilahyar/go-clean/internal/domain/user/handler"
)

type Application struct {
	Config *config.Config
	DB     *sql.DB

	User *userHandler.UserHandler

	Authorize *AuthorizeHandler
}

type AuthorizeHandler struct {
	Role       *roleHandler.RoleHandler
	Permission *permissionHandler.PermissionHandler

	Assignment *assignmentHandler.AssignmentHandler
}
