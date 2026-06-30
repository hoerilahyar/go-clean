package mapper

import (
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
)

// ToMeResponse maps the authenticated user information to a response DTO.
func ToMeResponse(
	me *entity.Me,
) *response.MeResponse {

	if me == nil {
		return nil
	}

	return &response.MeResponse{
		User:        ToUserResponse(me.User),
		Roles:       ToRoleResponses(me.Roles),
		Permissions: ToPermissionResponses(me.Permissions),
	}
}
