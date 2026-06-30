package repository

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/menu/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/menu/entity"
)

type MenuRepository interface {
	FindAll(ctx context.Context, req request.GetMenusRequest) ([]entity.Menu, error)

	FindByID(ctx context.Context, id uint64) (*entity.Menu, error)

	FindBySlug(ctx context.Context, slug string) (*entity.Menu, error)

	FindByIDs(ctx context.Context, ids []uint64) ([]entity.Menu, error)

	IsMenuExists(ctx context.Context, id uint64) (bool, error)

	IsSlugExists(ctx context.Context, slug string, excludeID uint64) (bool, error)

	IsNameExists(ctx context.Context, name string, excludeID uint64) (bool, error)

	Create(ctx context.Context, menu *entity.Menu) error

	Update(ctx context.Context, menu *entity.Menu) error

	Delete(ctx context.Context, id uint64, deletedBy uint64) error
}
