package mapper

import (
	roleEntity "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/entity"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/response"
)

// ToRoleResponse maps a role entity to a response DTO.
func ToRoleResponse(
	role roleEntity.Role,
) response.RoleResponse {

	return response.RoleResponse{
		ID:   role.ID,
		Name: role.Name,
		Slug: role.Slug,
	}
}

// ToRoleResponses maps role entities to response DTOs.
func ToRoleResponses(
	roles []roleEntity.Role,
) []response.RoleResponse {

	if len(roles) == 0 {
		return []response.RoleResponse{}
	}

	responses := make(
		[]response.RoleResponse,
		0,
		len(roles),
	)

	for _, role := range roles {

		responses = append(
			responses,
			ToRoleResponse(role),
		)
	}

	return responses
}
