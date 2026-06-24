package bootstrap

import "github.com/hoerilahyar/go-clean/internal/handler/http"

type Handlers struct {
	User       *http.UserHandler
	Role       *http.RoleHandler
	Permission *http.PermissionHandler
}

func NewHandlers(uc *Usecases) *Handlers {
	return &Handlers{
		User:       http.NewUserHandler(uc.User),
		Role:       http.NewRoleHandler(uc.Role),
		Permission: http.NewPermissionHandler(uc.Permission),
	}
}
