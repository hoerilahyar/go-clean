package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/repository"
)

type permissionUsecase struct {
	repo repository.PermissionRepository
}

func NewPermissionUsecase(repo repository.PermissionRepository) PermissionUsecase {
	return &permissionUsecase{
		repo: repo,
	}
}

func (u *permissionUsecase) GetAll(ctx context.Context, filter request.PermissionFilter) ([]entity.Permission, error) {
	return u.repo.FindAll(ctx, filter)
}

// func (u *permissionUsecase) GetByID(ctx context.Context, id uint64) (*entity.Permission, error) {
// 	if id == 0 {
// 		return nil, errors.New("invalid permission id")
// 	}

// 	permission, err := u.repo.FindByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if permission == nil {
// 		return nil, errors.New("permission not found")
// 	}

// 	return permission, nil
// }

func (u *permissionUsecase) GetByGroup(ctx context.Context, groupName string) ([]entity.Permission, error) {

	groupName = strings.TrimSpace(groupName)

	if groupName == "" {
		return nil, errors.New("group name is required")
	}

	return u.repo.FindByGroup(ctx, groupName)
}

func (u *permissionUsecase) Create(ctx context.Context, permission *entity.Permission) error {
	if permission == nil {
		return errors.New("permission is required")
	}

	permission.Name = strings.TrimSpace(permission.Name)
	permission.Slug = strings.TrimSpace(permission.Slug)
	permission.GroupName = strings.TrimSpace(permission.GroupName)

	if permission.Name == "" {
		return errors.New("permission name is required")
	}

	if permission.Slug == "" {
		return errors.New("permission slug is required")
	}

	if permission.GroupName == "" {
		return errors.New("permission group is required")
	}

	existing, err := u.repo.IsSlugExists(ctx, permission.Slug, 0)
	if err != nil {
		return err
	}

	if existing {
		return errors.New("permission slug already exists")
	}

	now := time.Now()

	permission.CreatedAt = now
	permission.UpdatedAt = now

	return u.repo.Create(ctx, permission)
}

func (u *permissionUsecase) Update(ctx context.Context, permission *entity.Permission) error {
	if permission == nil {
		return errors.New("permission is required")
	}

	if permission.ID == 0 {
		return errors.New("invalid permission id")
	}

	permission.Name = strings.TrimSpace(permission.Name)
	permission.Slug = strings.TrimSpace(permission.Slug)
	permission.GroupName = strings.TrimSpace(permission.GroupName)

	if permission.Name == "" {
		return errors.New("permission name is required")
	}

	if permission.Slug == "" {
		return errors.New("permission slug is required")
	}

	if permission.GroupName == "" {
		return errors.New("permission group is required")
	}

	current, err := u.repo.FindByID(ctx, permission.ID)
	if err != nil {
		return err
	}

	if current == nil {
		return errors.New("permission not found")
	}

	existing, err := u.repo.FindBySlug(ctx, permission.Slug)
	if err != nil {
		return err
	}

	if existing != nil && existing.ID != permission.ID {
		return errors.New("permission slug already exists")
	}

	permission.CreatedAt = current.CreatedAt
	permission.UpdatedAt = time.Now()

	return u.repo.Update(ctx, permission)
}

func (u *permissionUsecase) Delete(ctx context.Context, id uint64, deletedBy uint64) error {
	if id == 0 {
		return errors.New("invalid permission id")
	}

	permission, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if permission == nil {
		return errors.New("permission not found")
	}

	return u.repo.Delete(ctx, id, deletedBy)
}
