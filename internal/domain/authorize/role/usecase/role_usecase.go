package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/repository"
)

type roleUsecase struct {
	repo repository.RoleRepository
}

func NewRoleUsecase(repo repository.RoleRepository) RoleUsecase {
	return &roleUsecase{
		repo: repo,
	}
}

func (u *roleUsecase) GetAll(ctx context.Context, req request.GetRolesRequest) ([]entity.Role, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Slug = strings.TrimSpace(req.Slug)

	if req.ID != nil && *req.ID == 0 {
		return nil, errors.New("invalid id")
	}

	return u.repo.FindAll(ctx, req)
}

// func (u *roleUsecase) GetByID(ctx context.Context, id uint64) (*entity.Role, error) {
// 	if id == 0 {
// 		return nil, errors.New("invalid role id")
// 	}

// 	role, err := u.repo.FindByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if role == nil {
// 		return nil, errors.New("role not found")
// 	}

// 	return role, nil
// }

func (u *roleUsecase) Create(ctx context.Context, role *entity.Role) error {
	if role == nil {
		return errors.New("role is required")
	}

	role.Name = strings.TrimSpace(role.Name)

	if role.Name == "" {
		return errors.New("role name is required")
	}

	if role.Slug == "" {
		role.Slug = strings.ToLower(strings.ReplaceAll(role.Name, " ", "-"))
	}

	exists, err := u.repo.IsNameExists(ctx, role.Name, 0)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("role name already exists")
	}

	existing, err := u.repo.IsSlugExists(ctx, role.Slug, 0)
	if err != nil {
		return err
	}

	if existing {
		return errors.New("role slug already exists")
	}

	now := time.Now()

	role.CreatedAt = now
	role.UpdatedAt = now

	return u.repo.Create(ctx, role)
}

func (u *roleUsecase) Update(ctx context.Context, role *entity.Role) error {
	if role == nil {
		return errors.New("role is required")
	}

	if role.ID == 0 {
		return errors.New("invalid role id")
	}

	role.Name = strings.TrimSpace(role.Name)

	if role.Name == "" {
		return errors.New("role name is required")
	}

	current, err := u.repo.FindByID(ctx, role.ID)
	if err != nil {
		return err
	}

	if current == nil {
		return errors.New("role not found")
	}

	existing, err := u.repo.FindBySlug(ctx, role.Slug)
	if err != nil {
		return err
	}

	if existing != nil && existing.ID != role.ID {
		return errors.New("role slug already exists")
	}

	role.CreatedAt = current.CreatedAt
	role.UpdatedAt = time.Now()

	return u.repo.Update(ctx, role)
}

func (u *roleUsecase) Delete(ctx context.Context, id uint64, deletedBy uint64) error {
	if id == 0 {
		return errors.New("invalid role id")
	}

	role, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if role == nil {
		return errors.New("role not found")
	}

	return u.repo.Delete(ctx, id, deletedBy)
}
