package usecase

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/repository"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

type permissionUsecase struct {
	repo repository.PermissionRepository
}

func NewPermissionUsecase(
	repo repository.PermissionRepository,
) PermissionUsecase {
	return &permissionUsecase{
		repo: repo,
	}
}

// normalizePermission trims all input fields.
func normalizePermission(permission *entity.Permission) {

	permission.Name = strings.TrimSpace(permission.Name)
	permission.Slug = strings.TrimSpace(permission.Slug)
	permission.GroupName = strings.TrimSpace(permission.GroupName)
}

func (u *permissionUsecase) GetAll(
	ctx context.Context,
	filter request.PermissionFilter,
) ([]entity.Permission, error) {

	return u.repo.FindAll(ctx, filter)
}

func (u *permissionUsecase) GetByID(
	ctx context.Context,
	id uint64,
) (*entity.Permission, error) {

	if id == 0 {
		return nil, apperror.BadRequest("Invalid permission ID")
	}

	permission, err := u.repo.FindByID(ctx, id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.NotFound("Permission not found")
		}

		return nil, err
	}

	return permission, nil
}

func (u *permissionUsecase) GetByGroup(
	ctx context.Context,
	groupName string,
) ([]entity.Permission, error) {

	groupName = strings.TrimSpace(groupName)

	if groupName == "" {
		return nil, apperror.BadRequest("Permission group is required")
	}

	return u.repo.FindByGroup(ctx, groupName)
}
func (u *permissionUsecase) Create(
	ctx context.Context,
	permission *entity.Permission,
) error {

	if permission == nil {
		return apperror.BadRequest("Permission is required")
	}

	normalizePermission(permission)

	if permission.Name == "" {
		return apperror.BadRequest("Permission name is required")
	}

	if permission.Slug == "" {
		return apperror.BadRequest("Permission slug is required")
	}

	if permission.GroupName == "" {
		return apperror.BadRequest("Permission group is required")
	}

	exists, err := u.repo.IsSlugExists(
		ctx,
		permission.Slug,
		0,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict("Permission slug already exists")
	}

	// Uncomment if permission name must also be unique.
	//
	// exists, err = u.repo.IsNameExists(
	// 	ctx,
	// 	permission.Name,
	// 	0,
	// )
	// if err != nil {
	// 	return err
	// }
	//
	// if exists {
	// 	return apperror.Conflict("Permission name already exists")
	// }

	return u.repo.Create(ctx, permission)
}
func (u *permissionUsecase) Update(
	ctx context.Context,
	permission *entity.Permission,
) error {

	if permission == nil {
		return apperror.BadRequest("Permission is required")
	}

	if permission.ID == 0 {
		return apperror.BadRequest("Invalid permission ID")
	}

	normalizePermission(permission)

	if permission.Name == "" {
		return apperror.BadRequest("Permission name is required")
	}

	if permission.Slug == "" {
		return apperror.BadRequest("Permission slug is required")
	}

	if permission.GroupName == "" {
		return apperror.BadRequest("Permission group is required")
	}

	current, err := u.repo.FindByID(ctx, permission.ID)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return apperror.NotFound("Permission not found")
		}

		return err
	}

	exists, err := u.repo.IsSlugExists(
		ctx,
		permission.Slug,
		permission.ID,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict("Permission slug already exists")
	}

	// Uncomment if permission name must also be unique.
	//
	// exists, err = u.repo.IsNameExists(
	// 	ctx,
	// 	permission.Name,
	// 	permission.ID,
	// )
	// if err != nil {
	// 	return err
	// }
	//
	// if exists {
	// 	return apperror.Conflict("Permission name already exists")
	// }

	permission.CreatedAt = current.CreatedAt

	return u.repo.Update(ctx, permission)
}

func (u *permissionUsecase) Delete(
	ctx context.Context,
	id uint64,
	deletedBy uint64,
) error {

	if id == 0 {
		return apperror.BadRequest("Invalid permission ID")
	}

	_, err := u.repo.FindByID(ctx, id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return apperror.NotFound("Permission not found")
		}

		return err
	}

	return u.repo.Delete(
		ctx,
		id,
		deletedBy,
	)
}
