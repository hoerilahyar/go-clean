package bootstrap

import (
	"database/sql"

	assignmentRepo "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/repository"
	authorizationRepo "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/repository"
	permissionRepo "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/repository"
	roleRepo "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/repository"
	userRepo "github.com/hoerilahyar/go-clean/internal/domain/user/repository"
)

type Repositories struct {
	User userRepo.UserRepository

	Role       roleRepo.RoleRepository
	Permission permissionRepo.PermissionRepository
	Assignment assignmentRepo.AssignmentRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User: userRepo.NewUserRepository(db),

		Role:       roleRepo.NewRoleRepository(db),
		Permission: permissionRepo.NewPermissionRepository(db),
		Assignment: authorizationRepo.NewAssignmentRepository(db),
	}
}
