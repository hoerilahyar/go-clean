package bootstrap

import (
	"database/sql"

	"github.com/hoerilahyar/go-clean/internal/config"
	permissionHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/handler"
	roleHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/handler"
	userHandler "github.com/hoerilahyar/go-clean/internal/domain/user/handler"
)

type Application struct {
	Config *config.Config
	DB     *sql.DB

	UserHandler           *userHandler.UserHandler
	RoleHandler           *roleHandler.RoleHandler
	PermissionHandler     *permissionHandler.PermissionHandler
	RolePermissionHandler *roleHandler.RolePermissionHandler
}
