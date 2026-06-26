package repository

import (
	"context"
	"database/sql"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
)

type AssignmentRepository interface {
	// Role Permission
	AssignRolePermissions(ctx context.Context, roleID uint64, permissionIDs []uint64) error
	ReplaceRolePermissions(ctx context.Context, roleID uint64, permissionIDs []uint64) error
	FindRolePermissionsByRoleID(ctx context.Context, roleID uint64) ([]entity.RolePermission, error)
	DeleteRolePermission(ctx context.Context, roleID uint64, permissionID uint64) error

	// User Role
	AssignUserRoles(ctx context.Context, userID uint64, roleIDs []uint64) error
	ReplaceUserRoles(ctx context.Context, userID uint64, roleIDs []uint64) error
	FindUserRolesByUserID(ctx context.Context, userID uint64) ([]entity.UserRole, error)
	DeleteUserRole(ctx context.Context, userID uint64, roleID uint64) error

	// User Permission
	AssignUserPermissions(ctx context.Context, userID uint64, permissionIDs []uint64) error
	ReplaceUserPermissions(ctx context.Context, userID uint64, permissionIDs []uint64) error
	FindUserPermissionsByUserID(ctx context.Context, userID uint64) ([]entity.UserPermission, error)
	DeleteUserPermission(ctx context.Context, userID uint64, permissionID uint64) error

	// Authorization
	GetMe(ctx context.Context, userID uint64) (*entity.Me, error)
}

type assignmentRepository struct {
	db *sql.DB
}

func NewAssignmentRepository(db *sql.DB) AssignmentRepository {
	return &assignmentRepository{
		db: db,
	}
}
