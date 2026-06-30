package usecase

import (
	"context"
	"slices"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func (u *assignmentUsecase) AssignUserRoles(
	ctx context.Context,
	userID uint64,
	req request.AssignUserRoleRequest,
) error {

	// Validate user ID.
	if userID == 0 {
		return apperror.BadRequest("Invalid user ID")
	}

	// Validate role IDs.
	if len(req.RoleIDs) == 0 {
		return apperror.BadRequest("Role IDs are required")
	}

	// Sort role IDs for deterministic processing.
	slices.Sort(req.RoleIDs)

	// Remove duplicate role IDs.
	req.RoleIDs = slices.Compact(req.RoleIDs)

	// TODO:
	// Validate user exists.
	// Validate all roles exist.

	// Assign roles to the user.
	return u.repository.AssignUserRoles(
		ctx,
		userID,
		req.RoleIDs,
	)
}

func (u *assignmentUsecase) ReplaceUserRoles(
	ctx context.Context,
	userID uint64,
	req request.UpdateUserRolesRequest,
) error {

	// Validate user ID.
	if userID == 0 {
		return apperror.BadRequest("Invalid user ID")
	}

	// Sort role IDs for deterministic processing.
	slices.Sort(req.RoleIDs)

	// Remove duplicate role IDs.
	req.RoleIDs = slices.Compact(req.RoleIDs)

	// TODO:
	// Validate user exists.
	// Validate all roles exist.

	// Replace all roles assigned to the user.
	return u.repository.ReplaceUserRoles(
		ctx,
		userID,
		req.RoleIDs,
	)
}

func (u *assignmentUsecase) GetUserRoles(
	ctx context.Context,
	userID uint64,
) ([]entity.UserRole, error) {

	// Validate user ID.
	if userID == 0 {
		return nil, apperror.BadRequest("Invalid user ID")
	}

	// Retrieve roles assigned to the user.
	return u.repository.FindUserRolesByUserID(
		ctx,
		userID,
	)
}

func (u *assignmentUsecase) DeleteUserRole(
	ctx context.Context,
	userID uint64,
	roleID uint64,
) error {

	// Validate user ID.
	if userID == 0 {
		return apperror.BadRequest("Invalid user ID")
	}

	// Validate role ID.
	if roleID == 0 {
		return apperror.BadRequest("Invalid role ID")
	}

	// TODO:
	// Validate user exists.
	// Validate role exists.

	// Remove role from the user.
	return u.repository.DeleteUserRole(
		ctx,
		userID,
		roleID,
	)
}
