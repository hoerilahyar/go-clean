package repository

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
)

type PermissionRepository interface {
	// Queries
	FindAll(ctx context.Context, filter request.PermissionFilter) ([]entity.Permission, error)
	FindByID(ctx context.Context, id uint64) (*entity.Permission, error)
	FindByIDs(ctx context.Context, ids []uint64) ([]entity.Permission, error)
	FindBySlug(ctx context.Context, slug string) (*entity.Permission, error)
	FindByGroup(ctx context.Context, group string) ([]entity.Permission, error)

	// Validation
	IsPermissionExists(ctx context.Context, id uint64) (bool, error)
	IsSlugExists(ctx context.Context, slug string, excludeID uint64) (bool, error)
	IsNameExists(ctx context.Context, name string, excludeID uint64) (bool, error)

	// Commands
	Create(ctx context.Context, permission *entity.Permission) error
	Update(ctx context.Context, permission *entity.Permission) error
	Delete(ctx context.Context, id uint64, deletedBy uint64) error
}
