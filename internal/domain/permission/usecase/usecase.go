package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/permission/entity"
)

type PermissionUsecase interface {
	GetAll(ctx context.Context) ([]*entity.Permission, error)
	GetByID(ctx context.Context, id uint64) (*entity.Permission, error)
	GetByGroup(ctx context.Context, groupName string) ([]*entity.Permission, error)
	Create(ctx context.Context, permission *entity.Permission) error
	Update(ctx context.Context, permission *entity.Permission) error
	Delete(ctx context.Context, id uint64, deletedBy uint64) error
}
