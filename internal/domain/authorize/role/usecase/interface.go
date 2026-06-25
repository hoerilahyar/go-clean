package usecase

import (
	"context"

	permissionEntity "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/entity"
)

type RoleUsecase interface {
	GetAll(ctx context.Context, req request.GetRolesRequest) ([]entity.Role, error)
	// GetByID(ctx context.Context, id uint64) (*entity.Role, error)
	Create(ctx context.Context, role *entity.Role) error
	Update(ctx context.Context, role *entity.Role) error
	Delete(ctx context.Context, id uint64, deletedBy uint64) error
}

type RolePermissionUsecase interface {
	GetPermissions(ctx context.Context, roleID uint64) ([]*permissionEntity.Permission, error)

	Assign(ctx context.Context, roleID uint64, permissionIDs []uint64) error

	Sync(ctx context.Context, roleID uint64, permissionIDs []uint64) error

	Remove(ctx context.Context, roleID uint64, permissionID uint64) error
}
