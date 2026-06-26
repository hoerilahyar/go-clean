package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
	assignmentRepo "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/repository"
)

type AssignmentUsecase interface {
	// Role Permission
	AssignRolePermissions(ctx context.Context, roleID uint64, req request.AssignRolePermissionRequest) error
	ReplaceRolePermissions(ctx context.Context, roleID uint64, req request.UpdateRolePermissionsRequest) error
	GetRolePermissions(ctx context.Context, roleID uint64) ([]entity.RolePermission, error)
	DeleteRolePermission(ctx context.Context, roleID uint64, permissionID uint64) error

	// User Role
	AssignUserRoles(ctx context.Context, userID uint64, req request.AssignUserRoleRequest) error
	ReplaceUserRoles(ctx context.Context, userID uint64, req request.UpdateUserRolesRequest) error
	GetUserRoles(ctx context.Context, userID uint64) ([]entity.UserRole, error)
	DeleteUserRole(ctx context.Context, userID uint64, roleID uint64) error

	// User Permission
	AssignUserPermissions(ctx context.Context, userID uint64, req request.AssignUserPermissionRequest) error
	ReplaceUserPermissions(ctx context.Context, userID uint64, req request.UpdateUserPermissionsRequest) error
	GetUserPermissions(ctx context.Context, userID uint64) ([]entity.UserPermission, error)
	DeleteUserPermission(ctx context.Context, userID uint64, permissionID uint64) error

	// Authorization
	GetMe(ctx context.Context, userID uint64) (*response.MeResponse, error)
}

type assignmentUsecase struct {
	repository assignmentRepo.AssignmentRepository
}

func NewAssignmentUsecase(
	repository assignmentRepo.AssignmentRepository,
) AssignmentUsecase {
	return &assignmentUsecase{
		repository: repository,
	}
}
