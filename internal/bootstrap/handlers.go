package bootstrap

import (
	assignmentHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/handler"
	permissionHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/handler"
	roleHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/handler"
	userHandler "github.com/hoerilahyar/go-clean/internal/domain/user/handler"
)

type Handlers struct {
	User       *userHandler.UserHandler
	Role       *roleHandler.RoleHandler
	Permission *permissionHandler.PermissionHandler
	Assignment *assignmentHandler.AssignmentHandler
}

func NewHandlers(uc *Usecases) *Handlers {
	return &Handlers{
		User:       userHandler.NewUserHandler(uc.User),
		Role:       roleHandler.NewRoleHandler(uc.Role),
		Permission: permissionHandler.NewPermissionHandler(uc.Permission),
		Assignment: assignmentHandler.NewAssignmentHandler(uc.Assignment),
	}
}
