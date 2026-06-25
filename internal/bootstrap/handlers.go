package bootstrap

import (
	permissionHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/handler"
	roleHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/handler"
	userHandler "github.com/hoerilahyar/go-clean/internal/domain/user/handler"
)

type Handlers struct {
	User       *userHandler.UserHandler
	Role       *roleHandler.RoleHandler
	Permission *permissionHandler.PermissionHandler
}

func NewHandlers(uc *Usecases) *Handlers {
	return &Handlers{
		User:       userHandler.NewUserHandler(uc.User),
		Role:       roleHandler.NewRoleHandler(uc.Role),
		Permission: permissionHandler.NewPermissionHandler(uc.Permission),
	}
}
