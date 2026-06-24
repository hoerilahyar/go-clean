package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/permission/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/permission/repository"
)

type permissionUsecase struct {
	repo repository.PermissionRepository
}

func NewPermissionUsecase(repo repository.PermissionRepository) PermissionUsecase {
	return &permissionUsecase{
		repo: repo,
	}
}

func (u *permissionUsecase) GetAll(ctx context.Context) ([]*entity.Permission, error) {
	return u.repo.FindAll(ctx)
}

func (u *permissionUsecase) GetByID(ctx context.Context, id uint64) (*entity.Permission, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *permissionUsecase) GetByGroup(ctx context.Context, groupName string) ([]*entity.Permission, error) {
	return u.repo.FindByGroup(ctx, groupName)
}

func (u *permissionUsecase) Create(ctx context.Context, permission *entity.Permission) error {
	return u.repo.Create(ctx, permission)
}

func (u *permissionUsecase) Update(ctx context.Context, permission *entity.Permission) error {
	return u.repo.Update(ctx, permission)
}

func (u *permissionUsecase) Delete(ctx context.Context, id uint64, deletedBy uint64) error {
	return u.repo.Delete(ctx, id, deletedBy)
}
