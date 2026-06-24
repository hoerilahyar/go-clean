package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/user/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/user/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/user/entity"
)

type UserUsecase interface {
	GetAll(ctx context.Context, req request.GetUsersRequest) (*response.GetUsersResponse, error)
	GetByID(ctx context.Context, id uint64) (*entity.User, error)
	Create(ctx context.Context, req request.CreateUserRequest) (*entity.User, error)
	Update(ctx context.Context, req request.UpdateUserRequest) (*entity.User, error)
	Delete(ctx context.Context, id uint64, deletedBy uint64) error
}
