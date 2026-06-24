package repository

import (
	"context"
	"database/sql"

	"github.com/hoerilahyar/go-clean/internal/domain/role/entity"
)

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) FindAll(ctx context.Context) ([]*entity.Role, error) {
	query := `
		SELECT
			id,
			name,
			slug,
			description,
			is_active,
			created_at,
			updated_at,
			deleted_at,
			created_by,
			updated_by,
			deleted_by
		FROM roles
		WHERE deleted_at IS NULL
		ORDER BY name
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*entity.Role

	for rows.Next() {
		role := &entity.Role{}

		if err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.Slug,
			&role.Description,
			&role.IsActive,
			&role.CreatedAt,
			&role.UpdatedAt,
			&role.DeletedAt,
			&role.CreatedBy,
			&role.UpdatedBy,
			&role.DeletedBy,
		); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) FindByID(ctx context.Context, id uint64) (*entity.Role, error) {
	query := `
		SELECT
			id,
			name,
			slug,
			description,
			is_active,
			created_at,
			updated_at,
			deleted_at,
			created_by,
			updated_by,
			deleted_by
		FROM roles
		WHERE id = ?
		AND deleted_at IS NULL
	`

	role := &entity.Role{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&role.ID,
		&role.Name,
		&role.Slug,
		&role.Description,
		&role.IsActive,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
		&role.CreatedBy,
		&role.UpdatedBy,
		&role.DeletedBy,
	)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *roleRepository) FindBySlug(ctx context.Context, slug string) (*entity.Role, error) {
	query := `
		SELECT
			id,
			name,
			slug,
			description,
			is_active,
			created_at,
			updated_at,
			deleted_at,
			created_by,
			updated_by,
			deleted_by
		FROM roles
		WHERE slug = ?
		AND deleted_at IS NULL
	`

	role := &entity.Role{}

	err := r.db.QueryRowContext(ctx, query, slug).Scan(
		&role.ID,
		&role.Name,
		&role.Slug,
		&role.Description,
		&role.IsActive,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
		&role.CreatedBy,
		&role.UpdatedBy,
		&role.DeletedBy,
	)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *roleRepository) Create(ctx context.Context, role *entity.Role) error {
	query := `
		INSERT INTO roles (
			name,
			slug,
			description,
			is_active,
			created_by
		) VALUES (?, ?, ?, ?, ?)
	`

	res, err := r.db.ExecContext(
		ctx,
		query,
		role.Name,
		role.Slug,
		role.Description,
		role.IsActive,
		role.CreatedBy,
	)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	role.ID = uint64(id)

	return nil
}

func (r *roleRepository) Update(ctx context.Context, role *entity.Role) error {
	query := `
		UPDATE roles
		SET
			name = ?,
			slug = ?,
			description = ?,
			is_active = ?,
			updated_by = ?
		WHERE id = ?
		AND deleted_at IS NULL
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		role.Name,
		role.Slug,
		role.Description,
		role.IsActive,
		role.UpdatedBy,
		role.ID,
	)

	return err
}

func (r *roleRepository) Delete(ctx context.Context, id uint64, deletedBy uint64) error {
	query := `
		UPDATE roles
		SET
			deleted_at = NOW(),
			deleted_by = ?
		WHERE id = ?
		AND deleted_at IS NULL
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		deletedBy,
		id,
	)

	return err
}
