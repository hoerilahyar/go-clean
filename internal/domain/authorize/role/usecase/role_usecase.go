package usecase

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/repository"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

type roleUsecase struct {
	repo repository.RoleRepository
}

func NewRoleUsecase(
	repo repository.RoleRepository,
) RoleUsecase {

	return &roleUsecase{
		repo: repo,
	}
}

// normalizeRole trims request fields.
func normalizeRole(
	role *entity.Role,
) {

	role.Name = strings.TrimSpace(role.Name)
	role.Slug = strings.TrimSpace(role.Slug)
	if role.Description != nil {
		*role.Description = strings.TrimSpace(*role.Description)
	}
}

func (u *roleUsecase) GetAll(
	ctx context.Context,
	req request.GetRolesRequest,
) ([]entity.Role, error) {

	req.Name = strings.TrimSpace(req.Name)
	req.Slug = strings.TrimSpace(req.Slug)

	if req.ID != nil && *req.ID == 0 {
		return nil, apperror.BadRequest("Invalid role ID")
	}

	return u.repo.FindAll(
		ctx,
		req,
	)
}

func (u *roleUsecase) GetByID(
	ctx context.Context,
	id uint64,
) (*entity.Role, error) {

	if id == 0 {
		return nil, apperror.BadRequest("Invalid role ID")
	}

	role, err := u.repo.FindByID(
		ctx,
		id,
	)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.NotFound("Role not found")
		}

		return nil, err
	}

	return role, nil
}
func (u *roleUsecase) Create(
	ctx context.Context,
	role *entity.Role,
) error {

	if role == nil {
		return apperror.BadRequest("Role is required")
	}

	normalizeRole(role)

	if role.Name == "" {
		return apperror.BadRequest("Role name is required")
	}

	// Generate slug from role name when it is empty.
	if role.Slug == "" {
		role.Slug = strings.ToLower(
			strings.ReplaceAll(
				role.Name,
				" ",
				"-",
			),
		)
	}

	exists, err := u.repo.IsNameExists(
		ctx,
		role.Name,
		0,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Role name already exists",
		)
	}

	exists, err = u.repo.IsSlugExists(
		ctx,
		role.Slug,
		0,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Role slug already exists",
		)
	}

	return u.repo.Create(
		ctx,
		role,
	)
}
func (u *roleUsecase) Update(
	ctx context.Context,
	role *entity.Role,
) error {

	if role == nil {
		return apperror.BadRequest("Role is required")
	}

	if role.ID == 0 {
		return apperror.BadRequest("Invalid role ID")
	}

	normalizeRole(role)

	if role.Name == "" {
		return apperror.BadRequest("Role name is required")
	}

	if role.Slug == "" {
		role.Slug = strings.ToLower(
			strings.ReplaceAll(
				role.Name,
				" ",
				"-",
			),
		)
	}

	current, err := u.repo.FindByID(
		ctx,
		role.ID,
	)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return apperror.NotFound("Role not found")
		}

		return err
	}

	exists, err := u.repo.IsNameExists(
		ctx,
		role.Name,
		role.ID,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Role name already exists",
		)
	}

	exists, err = u.repo.IsSlugExists(
		ctx,
		role.Slug,
		role.ID,
	)
	if err != nil {
		return err
	}

	if exists {
		return apperror.Conflict(
			"Role slug already exists",
		)
	}

	role.CreatedAt = current.CreatedAt
	role.CreatedBy = current.CreatedBy

	return u.repo.Update(
		ctx,
		role,
	)
}

func (u *roleUsecase) Delete(
	ctx context.Context,
	id uint64,
	deletedBy uint64,
) error {

	if id == 0 {
		return apperror.BadRequest("Invalid role ID")
	}

	_, err := u.repo.FindByID(
		ctx,
		id,
	)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return apperror.NotFound("Role not found")
		}

		return err
	}

	return u.repo.Delete(
		ctx,
		id,
		deletedBy,
	)
}
