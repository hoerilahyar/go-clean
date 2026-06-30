package mapper

import (
	permissionEntity "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/response"
)

// ToPermissionResponse maps a permission entity to a response DTO.
func ToPermissionResponse(
	permission permissionEntity.Permission,
) response.PermissionResponse {

	return response.PermissionResponse{
		ID:    permission.ID,
		Name:  permission.Name,
		Slug:  permission.Slug,
		Group: permission.GroupName,
	}
}

// ToPermissionResponses maps permission entities to response DTOs.
func ToPermissionResponses(
	permissions []permissionEntity.Permission,
) []response.PermissionResponse {

	if len(permissions) == 0 {
		return []response.PermissionResponse{}
	}

	responses := make(
		[]response.PermissionResponse,
		0,
		len(permissions),
	)

	for _, permission := range permissions {

		responses = append(
			responses,
			ToPermissionResponse(permission),
		)
	}

	return responses
}
