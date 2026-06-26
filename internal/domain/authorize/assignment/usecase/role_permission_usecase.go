package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
)

func (u *assignmentUsecase) AssignRolePermissions(
	ctx context.Context,
	roleID uint64,
	req request.AssignRolePermissionRequest,
) error {

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

	return u.repository.DeleteRolePermission(
		ctx,
		roleID,
		permissionID,
	)
}
