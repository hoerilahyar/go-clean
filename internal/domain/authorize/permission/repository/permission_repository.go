package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
)

type permissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) PermissionRepository {
	return &permissionRepository{
		db: db,
	}
}

func (r *permissionRepository) FindAll(ctx context.Context, filter request.PermissionFilter) ([]entity.Permission, error) {
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
	`

	args := make([]any, 0)

	if filter.ID != nil {
		query += " AND id = ?"
		args = append(args, *filter.ID)
	}

	if filter.Name != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+filter.Name+"%")
	}

	if filter.Slug != "" {
		query += " AND slug LIKE ?"
		args = append(args, "%"+filter.Slug+"%")
	}

	if filter.GroupName != "" {
		query += " AND group_name = ?"
		args = append(args, filter.GroupName)
	}

	query += " ORDER BY group_name ASC, name ASC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := make([]entity.Permission, 0)

	for rows.Next() {
		var p entity.Permission

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

func (r *permissionRepository) FindByIDs(ctx context.Context, ids []uint64) ([]entity.Permission, error) {

	if len(ids) == 0 {
		return []entity.Permission{}, nil
	}

	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))

	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT
			id,
			name,
			slug,
			group_name,
			description,
			created_at,
			updated_at
		FROM permissions
		WHERE id IN (%s)
		AND deleted_at IS NULL
		ORDER BY id
	`, strings.Join(placeholders, ","))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []entity.Permission

	for rows.Next() {
		var permission entity.Permission

		if err := rows.Scan(
			&permission.ID,
			&permission.Name,
			&permission.Slug,
			&permission.GroupName,
			&permission.Description,
			&permission.CreatedAt,
			&permission.UpdatedAt,
		); err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
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

func (r *permissionRepository) FindByGroup(ctx context.Context, group string) ([]entity.Permission, error) {
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

	rows, err := r.db.QueryContext(ctx, query, group)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []entity.Permission

	for rows.Next() {
		var p entity.Permission

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

func (r *permissionRepository) IsNameExists(
	ctx context.Context,
	name string,
	excludeID uint64,
) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM permissions
			WHERE name = ?
			  AND deleted_at IS NULL
			  AND (? = 0 OR id <> ?)
		)
	`

	var exists bool

	err := r.db.QueryRowContext(
		ctx,
		query,
		name,
		excludeID,
		excludeID,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *permissionRepository) IsPermissionExists(ctx context.Context, id uint64) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM permissions
			WHERE id = ?
			  AND deleted_at IS NULL
		)
	`

	var exists bool

	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *permissionRepository) IsSlugExists(ctx context.Context, slug string, excludeID uint64) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM permissions
			WHERE slug = ?
			  AND deleted_at IS NULL
			  AND (? = 0 OR id <> ?)
		)
	`

	var exists bool

	err := r.db.QueryRowContext(
		ctx,
		query,
		slug,
		excludeID,
		excludeID,
	).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
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
