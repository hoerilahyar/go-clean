package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/menu/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/menu/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/menu/entity"
)

type MenuUsecase interface {
	GetAll(ctx context.Context, req request.GetMenusRequest) (*response.GetMenusResponse, error)

	GetByID(ctx context.Context, id uint64) (*entity.Menu, error)

	Create(ctx context.Context, req request.CreateMenuRequest) (*entity.Menu, error)

	Update(ctx context.Context, req request.UpdateMenuRequest) (*entity.Menu, error)

	Delete(ctx context.Context, id uint64, deletedBy uint64) error
}
