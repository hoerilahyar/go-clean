package bootstrap

import (
	"database/sql"

	permissionRepo "github.com/hoerilahyar/go-clean/internal/domain/permission/repository"
	roleRepo "github.com/hoerilahyar/go-clean/internal/domain/role/repository"
	userRepo "github.com/hoerilahyar/go-clean/internal/domain/user/repository"
)

type Repositories struct {
	User       userRepo.UserRepository
	Role       roleRepo.RoleRepository
	Permission permissionRepo.PermissionRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User:       userRepo.NewUserRepository(db),
		Role:       roleRepo.NewRoleRepository(db),
		Permission: permissionRepo.NewPermissionRepository(db),
	}
}
