package bootstrap

import (
	authNHandler "github.com/hoerilahyar/go-clean/internal/domain/authentication/handler"
	assignmentHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/handler"
	permissionHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/handler"
	roleHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/handler"
	menuHandler "github.com/hoerilahyar/go-clean/internal/domain/menu/handler"
	userHandler "github.com/hoerilahyar/go-clean/internal/domain/user/handler"
)

type Handlers struct {
	User       *userHandler.UserHandler
	Role       *roleHandler.RoleHandler
	Permission *permissionHandler.PermissionHandler
	Assignment *assignmentHandler.AssignmentHandler
	AuthN      *authNHandler.AuthenticationHandler
	Menu       *menuHandler.MenuHandler
}

func NewHandlers(uc *Usecases) *Handlers {
	return &Handlers{
		User:       userHandler.NewUserHandler(uc.User),
		Role:       roleHandler.NewRoleHandler(uc.Role),
		Permission: permissionHandler.NewPermissionHandler(uc.Permission),
		Assignment: assignmentHandler.NewAssignmentHandler(uc.Assignment),
		AuthN:      authNHandler.NewAuthenticationHandler(uc.AuthN),
		Menu:       menuHandler.NewMenuHandler(uc.Menu),
	}
}
