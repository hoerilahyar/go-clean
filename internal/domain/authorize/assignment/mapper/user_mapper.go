package mapper

import (
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/user/entity"
)

// ToUserResponse maps a user entity to a response DTO.
func ToUserResponse(
	user entity.User,
) response.UserResponse {

	return response.UserResponse{
		ID:       user.ID,
		Name:     user.FullName,
		Username: user.Username,
		Email:    user.Email,
		Status:   user.Status,
	}
}
