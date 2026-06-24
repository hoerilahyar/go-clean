package repository

import (
	"context"
	"database/sql"

	"github.com/hoerilahyar/go-clean/internal/domain/permission/entity"
)

type permissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) PermissionRepository {
	return &permissionRepository{
		db: db,
	}
}

func (r *permissionRepository) FindAll(ctx context.Context) ([]*entity.Permission, error) {
	query := `
		SELECT
			id,
			name,
			slug,
			group_name,
			description,
			created_at,
			updated_at,
			deleted_at,
			created_by,
			updated_by,
			deleted_by
		FROM permissions
		WHERE deleted_at IS NULL
		ORDER BY group_name, name
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*entity.Permission

	for rows.Next() {
		p := &entity.Permission{}

		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Slug,
			&p.GroupName,
			&p.Description,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.DeletedAt,
			&p.CreatedBy,
			&p.UpdatedBy,
			&p.DeletedBy,
		); err != nil {
			return nil, err
		}

		permissions = append(permissions, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *permissionRepository) FindByID(ctx context.Context, id uint64) (*entity.Permission, error) {
	query := `
		SELECT
			id,
			name,
			slug,
			group_name,
			description,
			created_at,
			updated_at,
			deleted_at,
			created_by,
			updated_by,
			deleted_by
		FROM permissions
		WHERE id = ?
		AND deleted_at IS NULL
	`

	p := &entity.Permission{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Slug,
		&p.GroupName,
		&p.Description,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DeletedAt,
		&p.CreatedBy,
		&p.UpdatedBy,
		&p.DeletedBy,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *permissionRepository) FindBySlug(ctx context.Context, slug string) (*entity.Permission, error) {
	query := `
		SELECT
			id,
			name,
			slug,
			group_name,
			description,
			created_at,
			updated_at,
			deleted_at,
			created_by,
			updated_by,
			deleted_by
		FROM permissions
		WHERE slug = ?
		AND deleted_at IS NULL
	`

	p := &entity.Permission{}

	err := r.db.QueryRowContext(ctx, query, slug).Scan(
		&p.ID,
		&p.Name,
		&p.Slug,
		&p.GroupName,
		&p.Description,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DeletedAt,
		&p.CreatedBy,
		&p.UpdatedBy,
		&p.DeletedBy,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *permissionRepository) FindByGroup(ctx context.Context, groupName string) ([]*entity.Permission, error) {
	query := `
		SELECT
			id,
			name,
			slug,
			group_name,
			description,
			created_at,
			updated_at,
			deleted_at,
			created_by,
			updated_by,
			deleted_by
		FROM permissions
		WHERE group_name = ?
		AND deleted_at IS NULL
		ORDER BY name
	`

	rows, err := r.db.QueryContext(ctx, query, groupName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*entity.Permission

	for rows.Next() {
		p := &entity.Permission{}

		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Slug,
			&p.GroupName,
			&p.Description,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.DeletedAt,
			&p.CreatedBy,
			&p.UpdatedBy,
			&p.DeletedBy,
		); err != nil {
			return nil, err
		}

		permissions = append(permissions, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *permissionRepository) Create(ctx context.Context, p *entity.Permission) error {
	query := `
		INSERT INTO permissions (
			name,
			slug,
			group_name,
			description,
			created_by
		) VALUES (?, ?, ?, ?, ?)
	`

	res, err := r.db.ExecContext(
		ctx,
		query,
		p.Name,
		p.Slug,
		p.GroupName,
		p.Description,
		p.CreatedBy,
	)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = uint64(id)

	return nil
}

func (r *permissionRepository) Update(ctx context.Context, p *entity.Permission) error {
	query := `
		UPDATE permissions
		SET
			name = ?,
			group_name = ?,
			description = ?,
			updated_by = ?
		WHERE id = ?
		AND deleted_at IS NULL
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		p.Name,
		p.GroupName,
		p.Description,
		p.UpdatedBy,
		p.ID,
	)

	return err
}

func (r *permissionRepository) Delete(ctx context.Context, id uint64, deletedBy uint64) error {
	query := `
		UPDATE permissions
		SET
			deleted_at = NOW(),
			deleted_by = ?
		WHERE id = ?
		AND deleted_at IS NULL
	`

	_, err := r.db.ExecContext(ctx, query, deletedBy, id)

	return err
}
