package usecase

import (
	"context"
	"errors"

	permissionEntity "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
	permissionRepository "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/repository"
	roleRepository "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/repository"
)

type rolePermissionUsecase struct {
	roleRepo           roleRepository.RoleRepository
	permissionRepo     permissionRepository.PermissionRepository
	rolePermissionRepo roleRepository.RolePermissionRepository
}

func NewRolePermissionUsecase(
	roleRepo roleRepository.RoleRepository,
	permissionRepo permissionRepository.PermissionRepository,
	rolePermissionRepo roleRepository.RolePermissionRepository,
) RolePermissionUsecase {
	return &rolePermissionUsecase{
		roleRepo:           roleRepo,
		permissionRepo:     permissionRepo,
		rolePermissionRepo: rolePermissionRepo,
	}
}

func (u *rolePermissionUsecase) GetPermissions(
	ctx context.Context,
	roleID uint64,
) ([]*permissionEntity.Permission, error) {

	if roleID == 0 {
		return nil, errors.New("invalid role id")
	}

	role, err := u.roleRepo.FindByID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	if role == nil {
		return nil, errors.New("role not found")
	}

	return u.rolePermissionRepo.GetPermissions(ctx, roleID)
}

func (u *rolePermissionUsecase) Assign(
	ctx context.Context,
	roleID uint64,
	permissionIDs []uint64,
) error {

	if roleID == 0 {
		return errors.New("invalid role id")
	}

	if len(permissionIDs) == 0 {
		return errors.New("permission ids are required")
	}

	role, err := u.roleRepo.FindByID(ctx, roleID)
	if err != nil {
		return err
	}

	if role == nil {
		return errors.New("role not found")
	}

	for _, permissionID := range permissionIDs {

		permission, err := u.permissionRepo.FindByID(ctx, permissionID)
		if err != nil {
			return err
		}

		if permission == nil {
			return errors.New("permission not found")
		}
	}

	return u.rolePermissionRepo.Assign(
		ctx,
		roleID,
		permissionIDs,
	)
}

func (u *rolePermissionUsecase) Sync(
	ctx context.Context,
	roleID uint64,
	permissionIDs []uint64,
) error {

	if roleID == 0 {
		return errors.New("invalid role id")
	}

	role, err := u.roleRepo.FindByID(ctx, roleID)
	if err != nil {
		return err
	}

	if role == nil {
		return errors.New("role not found")
	}

	for _, permissionID := range permissionIDs {

		permission, err := u.permissionRepo.FindByID(ctx, permissionID)
		if err != nil {
			return err
		}

		if permission == nil {
			return errors.New("permission not found")
		}
	}

	return u.rolePermissionRepo.Sync(
		ctx,
		roleID,
		permissionIDs,
	)
}

func (u *rolePermissionUsecase) Remove(
	ctx context.Context,
	roleID uint64,
	permissionID uint64,
) error {

	if roleID == 0 {
		return errors.New("invalid role id")
	}

	if permissionID == 0 {
		return errors.New("invalid permission id")
	}

	role, err := u.roleRepo.FindByID(ctx, roleID)
	if err != nil {
		return err
	}

	if role == nil {
		return errors.New("role not found")
	}

	permission, err := u.permissionRepo.FindByID(ctx, permissionID)
	if err != nil {
		return err
	}

	if permission == nil {
		return errors.New("permission not found")
	}

	return u.rolePermissionRepo.Remove(
		ctx,
		roleID,
		permissionID,
	)
}
