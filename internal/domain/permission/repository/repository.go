package repository

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/permission/entity"
)

type PermissionRepository interface {
	FindAll(ctx context.Context) ([]*entity.Permission, error)
	FindByID(ctx context.Context, id uint64) (*entity.Permission, error)
	FindBySlug(ctx context.Context, slug string) (*entity.Permission, error)
	FindByGroup(ctx context.Context, groupName string) ([]*entity.Permission, error)
	Create(ctx context.Context, permission *entity.Permission) error
	Update(ctx context.Context, permission *entity.Permission) error
	Delete(ctx context.Context, id uint64, deletedBy uint64) error
}
