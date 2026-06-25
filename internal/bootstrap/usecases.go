package bootstrap

import (
	permissionUsecase "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/usecase"
	roleUsecase "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/usecase"
	userUsecase "github.com/hoerilahyar/go-clean/internal/domain/user/usecase"
)

type Usecases struct {
	User       userUsecase.UserUsecase
	Role       roleUsecase.RoleUsecase
	Permission permissionUsecase.PermissionUsecase
}

func NewUsecases(repo *Repositories) *Usecases {
	return &Usecases{
		User:       userUsecase.NewUserUsecase(repo.User),
		Role:       roleUsecase.NewRoleUsecase(repo.Role),
		Permission: permissionUsecase.NewPermissionUsecase(repo.Permission),
	}
}
