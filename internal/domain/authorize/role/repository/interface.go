package repository

import (
	"context"

	permissionEntity "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/entity"
)

type RoleRepository interface {
	FindAll(ctx context.Context, req request.GetRolesRequest) ([]entity.Role, error)
	FindByID(ctx context.Context, id uint64) (*entity.Role, error)
	FindByIDs(ctx context.Context, ids []uint64) ([]*entity.Role, error)
	FindBySlug(ctx context.Context, slug string) (*entity.Role, error)

	IsSlugExists(ctx context.Context, slug string, excludeID uint64) (bool, error)
	IsNameExists(ctx context.Context, name string, excludeID uint64) (bool, error)
	IsRoleExists(ctx context.Context, id uint64) (bool, error)

	Create(ctx context.Context, role *entity.Role) error
	Update(ctx context.Context, role *entity.Role) error
	Delete(ctx context.Context, id uint64, deletedBy uint64) error
}

type RolePermissionRepository interface {
	GetPermissions(ctx context.Context, roleID uint64) ([]*permissionEntity.Permission, error)

	Assign(ctx context.Context, roleID uint64, permissionIDs []uint64) error

	Remove(ctx context.Context, roleID uint64, permissionID uint64) error

	Sync(ctx context.Context, roleID uint64, permissionIDs []uint64) error
}
