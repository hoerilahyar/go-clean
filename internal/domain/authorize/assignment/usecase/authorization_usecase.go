package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/mapper"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func (u *assignmentUsecase) GetMe(
	ctx context.Context,
	userID uint64,
) (*response.MeResponse, error) {

	// Validate user ID.
	if userID == 0 {
		return nil, apperror.BadRequest("Invalid user ID")
	}

	// Retrieve authenticated user information.
	me, err := u.repository.GetMe(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	// Ensure the user exists.
	if me == nil {
		return nil, apperror.ErrUserNotFound
	}

	// Map entity to response DTO.
	return mapper.ToMeResponse(me), nil
}
