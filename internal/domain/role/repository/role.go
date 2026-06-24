package repository

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/role/entity"
)

type RoleRepository interface {
	FindAll(ctx context.Context) ([]*entity.Role, error)
	FindByID(ctx context.Context, id uint64) (*entity.Role, error)
	FindBySlug(ctx context.Context, slug string) (*entity.Role, error)
	Create(ctx context.Context, role *entity.Role) error
	Update(ctx context.Context, role *entity.Role) error
	Delete(ctx context.Context, id uint64, deletedBy uint64) error
}
