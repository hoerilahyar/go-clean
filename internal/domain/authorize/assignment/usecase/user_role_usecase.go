package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
)

func (u *assignmentUsecase) AssignUserRoles(
	ctx context.Context,
	userID uint64,
	req request.AssignUserRoleRequest,
) error {

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

	return u.repository.DeleteUserRole(
		ctx,
		userID,
		roleID,
	)
}
