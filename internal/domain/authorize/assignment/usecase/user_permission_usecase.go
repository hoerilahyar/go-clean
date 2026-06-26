package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
)

func (u *assignmentUsecase) AssignUserPermissions(
	ctx context.Context,
	userID uint64,
	req request.AssignUserPermissionRequest,
) error {

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

	return u.repository.DeleteUserPermission(
		ctx,
		userID,
		permissionID,
	)
}
