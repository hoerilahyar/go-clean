package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/role/entity"
)

type RoleUsecase interface {
	GetAll(ctx context.Context) ([]*entity.Role, error)
	GetByID(ctx context.Context, id uint64) (*entity.Role, error)
	Create(ctx context.Context, role *entity.Role) error
	Update(ctx context.Context, role *entity.Role) error
	Delete(ctx context.Context, id uint64, deletedBy uint64) error
}
