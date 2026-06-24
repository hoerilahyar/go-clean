package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/role/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/role/repository"
)

type roleUsecase struct {
	repo repository.RoleRepository
}

func NewRoleUsecase(repo repository.RoleRepository) RoleUsecase {
	return &roleUsecase{
		repo: repo,
	}
}

func (u *roleUsecase) GetAll(ctx context.Context) ([]*entity.Role, error) {
	return u.repo.FindAll(ctx)
}

func (u *roleUsecase) GetByID(ctx context.Context, id uint64) (*entity.Role, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *roleUsecase) Create(ctx context.Context, role *entity.Role) error {
	return u.repo.Create(ctx, role)
}

func (u *roleUsecase) Update(ctx context.Context, role *entity.Role) error {
	return u.repo.Update(ctx, role)
}

func (u *roleUsecase) Delete(ctx context.Context, id uint64, deletedBy uint64) error {
	return u.repo.Delete(ctx, id, deletedBy)
}
