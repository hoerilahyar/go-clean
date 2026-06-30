package usecase

import (
	"context"
	"slices"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func (u *assignmentUsecase) AssignUserPermissions(
	ctx context.Context,
	userID uint64,
	req request.AssignUserPermissionRequest,
) error {

	// Validate user ID.
	if userID == 0 {
		return apperror.BadRequest("Invalid user ID")
	}

	// Validate permission IDs.
	if len(req.PermissionIDs) == 0 {
		return apperror.BadRequest("Permission IDs are required")
	}

	// Sort permission IDs for deterministic processing.
	slices.Sort(req.PermissionIDs)

	// Remove duplicate permission IDs.
	req.PermissionIDs = slices.Compact(req.PermissionIDs)

	// TODO:
	// Validate user exists.
	// Validate all permissions exist.

	// Assign permissions to the user.
	return u.repository.AssignUserPermissions(
		ctx,
		userID,
		req.PermissionIDs,
	)
}

func (u *assignmentUsecase) ReplaceUserPermissions(
	ctx context.Context,
	userID uint64,
	req request.UpdateUserPermissionsRequest,
) error {

	// Validate user ID.
	if userID == 0 {
		return apperror.BadRequest("Invalid user ID")
	}

	// Sort permission IDs for deterministic processing.
	slices.Sort(req.PermissionIDs)

	// Remove duplicate permission IDs.
	req.PermissionIDs = slices.Compact(req.PermissionIDs)

	// TODO:
	// Validate user exists.
	// Validate all permissions exist.

	// Replace all permissions assigned to the user.
	return u.repository.ReplaceUserPermissions(
		ctx,
		userID,
		req.PermissionIDs,
	)
}

func (u *assignmentUsecase) GetUserPermissions(
	ctx context.Context,
	userID uint64,
) ([]entity.UserPermission, error) {

	// Validate user ID.
	if userID == 0 {
		return nil, apperror.BadRequest("Invalid user ID")
	}

	// Retrieve permissions assigned to the user.
	return u.repository.FindUserPermissionsByUserID(
		ctx,
		userID,
	)
}

func (u *assignmentUsecase) DeleteUserPermission(
	ctx context.Context,
	userID uint64,
	permissionID uint64,
) error {

	// Validate user ID.
	if userID == 0 {
		return apperror.BadRequest("Invalid user ID")
	}

	// Validate permission ID.
	if permissionID == 0 {
		return apperror.BadRequest("Invalid permission ID")
	}

	// TODO:
	// Validate user exists.
	// Validate permission exists.

	// Remove permission from the user.
	return u.repository.DeleteUserPermission(
		ctx,
		userID,
		permissionID,
	)
}
