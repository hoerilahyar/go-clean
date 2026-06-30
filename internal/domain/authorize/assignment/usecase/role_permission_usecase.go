package usecase

import (
	"context"
	"slices"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func (u *assignmentUsecase) AssignRolePermissions(
	ctx context.Context,
	roleID uint64,
	req request.AssignRolePermissionRequest,
) error {

	// Validate role ID.
	if roleID == 0 {
		return apperror.BadRequest("Invalid role ID")
	}

	// Validate permission IDs.
	if len(req.PermissionIDs) == 0 {
		return apperror.BadRequest("Permission IDs are required")
	}

	// Remove duplicate permission IDs.
	req.PermissionIDs = slices.Compact(req.PermissionIDs)

	// Sort permission IDs for deterministic processing.
	slices.Sort(req.PermissionIDs)

	// TODO:
	// Validate role exists.
	// Validate all permissions exist.

	// Assign permissions to the role.
	return u.repository.AssignRolePermissions(
		ctx,
		roleID,
		req.PermissionIDs,
	)
}

func (u *assignmentUsecase) ReplaceRolePermissions(
	ctx context.Context,
	roleID uint64,
	req request.UpdateRolePermissionsRequest,
) error {

	// Validate role ID.
	if roleID == 0 {
		return apperror.BadRequest("Invalid role ID")
	}

	// Remove duplicate permission IDs.
	req.PermissionIDs = slices.Compact(req.PermissionIDs)

	// Sort permission IDs for deterministic processing.
	slices.Sort(req.PermissionIDs)

	// TODO:
	// Validate role exists.
	// Validate all permissions exist.

	// Replace all permissions assigned to the role.
	return u.repository.ReplaceRolePermissions(
		ctx,
		roleID,
		req.PermissionIDs,
	)
}

func (u *assignmentUsecase) GetRolePermissions(
	ctx context.Context,
	roleID uint64,
) ([]entity.RolePermission, error) {

	// Validate role ID.
	if roleID == 0 {
		return nil, apperror.BadRequest("Invalid role ID")
	}

	// Retrieve permissions assigned to the role.
	return u.repository.FindRolePermissionsByRoleID(
		ctx,
		roleID,
	)
}

func (u *assignmentUsecase) DeleteRolePermission(
	ctx context.Context,
	roleID uint64,
	permissionID uint64,
) error {

	// Validate role ID.
	if roleID == 0 {
		return apperror.BadRequest("Invalid role ID")
	}

	// Validate permission ID.
	if permissionID == 0 {
		return apperror.BadRequest("Invalid permission ID")
	}

	// TODO:
	// Validate role exists.
	// Validate permission exists.

	// Remove permission from the role.
	return u.repository.DeleteRolePermission(
		ctx,
		roleID,
		permissionID,
	)
}
