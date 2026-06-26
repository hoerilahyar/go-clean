package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/response"
)

func (u *assignmentUsecase) GetMe(
	ctx context.Context,
	userID uint64,
) (*response.MeResponse, error) {

	me, err := u.repository.GetMe(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := &response.MeResponse{
		User: response.UserResponse{
			ID:       me.User.ID,
			Name:     me.User.FullName,
			Username: me.User.Username,
			Email:    me.User.Email,
			Status:   me.User.Status,
		},
	}

	for _, role := range me.Roles {
		result.Roles = append(result.Roles, response.RoleResponse{
			ID:   role.ID,
			Name: role.Name,
			Slug: role.Slug,
		})
	}

	for _, permission := range me.Permissions {
		result.Permissions = append(result.Permissions, response.PermissionResponse{
			ID:    permission.ID,
			Name:  permission.Name,
			Slug:  permission.Slug,
			Group: permission.GroupName,
		})
	}

	return result, nil
}
