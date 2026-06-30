package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

const permissionColumns = `
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
`

type permissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) PermissionRepository {
	return &permissionRepository{
		db: db,
	}
}

// scanPermission scans a permission from sql.Row or sql.Rows.
func scanPermission(scanner interface {
	Scan(dest ...any) error
}) (*entity.Permission, error) {

	p := &entity.Permission{}

	err := scanner.Scan(
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

// findOne executes a query that returns a single permission.
func (r *permissionRepository) findOne(
	ctx context.Context,
	query string,
	args ...any,
) (*entity.Permission, error) {

	row := r.db.QueryRowContext(ctx, query, args...)

	permission, err := scanPermission(row)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, apperror.Internal("Failed to scan permission", err)
	}

	return permission, nil
}

// exists executes an EXISTS query.
func (r *permissionRepository) exists(
	ctx context.Context,
	query string,
	args ...any,
) (bool, error) {

	var exists bool

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return false, err
		}

		return false, apperror.Internal(
			"Failed to retrieve permissions",
			err,
		)
	}

	return exists, nil
}

func (r *permissionRepository) FindAll(
	ctx context.Context,
	filter request.PermissionFilter,
) ([]entity.Permission, error) {

	query := `
		SELECT
			` + permissionColumns + `
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
		return nil, apperror.Internal(
			"Failed to retrieve permissions",
			err,
		)
	}

	defer rows.Close()

	permissions := make([]entity.Permission, 0)

	for rows.Next() {

		permission, err := scanPermission(rows)
		if err != nil {
			return nil, apperror.Internal(
				"Failed to scan permission",
				err,
			)
		}

		permissions = append(permissions, *permission)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.Internal(
			"Failed to iterate permissions",
			err,
		)
	}

	return permissions, nil
}

func (r *permissionRepository) FindByID(
	ctx context.Context,
	id uint64,
) (*entity.Permission, error) {

	query := `
		SELECT
			` + permissionColumns + `
		FROM permissions
		WHERE id = ?
		  AND deleted_at IS NULL
	`

	return r.findOne(ctx, query, id)
}

func (r *permissionRepository) FindBySlug(
	ctx context.Context,
	slug string,
) (*entity.Permission, error) {

	query := `
		SELECT
			` + permissionColumns + `
		FROM permissions
		WHERE slug = ?
		  AND deleted_at IS NULL
	`

	return r.findOne(ctx, query, slug)
}

// buildPlaceholders builds placeholders for IN queries.
func buildPlaceholders(count int) string {

	placeholders := make([]string, count)

	for i := range placeholders {
		placeholders[i] = "?"
	}

	return strings.Join(placeholders, ",")
}

func (r *permissionRepository) FindByIDs(
	ctx context.Context,
	ids []uint64,
) ([]entity.Permission, error) {

	if len(ids) == 0 {
		return []entity.Permission{}, nil
	}

	args := make([]any, len(ids))

	for i, id := range ids {
		args[i] = id
	}

	query := `
		SELECT
			id,
			name,
			slug,
			group_name,
			description,
			created_at,
			updated_at
		FROM permissions
		WHERE id IN (` + buildPlaceholders(len(ids)) + `)
		  AND deleted_at IS NULL
		ORDER BY id
	`

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, apperror.Internal(
			"Failed to retrieve permissions",
			err,
		)
	}

	defer rows.Close()

	permissions := make([]entity.Permission, 0, len(ids))

	for rows.Next() {

		var permission entity.Permission

		err := rows.Scan(
			&permission.ID,
			&permission.Name,
			&permission.Slug,
			&permission.GroupName,
			&permission.Description,
			&permission.CreatedAt,
			&permission.UpdatedAt,
		)
		if err != nil {
			return nil, apperror.Internal(
				"Failed to scan permission",
				err,
			)
		}

		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.Internal(
			"Failed to iterate permissions",
			err,
		)
	}

	return permissions, nil
}

func (r *permissionRepository) FindByGroup(
	ctx context.Context,
	group string,
) ([]entity.Permission, error) {

	query := `
		SELECT
			` + permissionColumns + `
		FROM permissions
		WHERE group_name = ?
		  AND deleted_at IS NULL
		ORDER BY name
	`

	rows, err := r.db.QueryContext(ctx, query, group)
	if err != nil {
		return nil, apperror.Internal(
			"Failed to retrieve permissions",
			err,
		)
	}

	defer rows.Close()

	permissions := make([]entity.Permission, 0)

	for rows.Next() {

		permission, err := scanPermission(rows)
		if err != nil {
			return nil, apperror.Internal(
				"Failed to scan permission",
				err,
			)
		}

		permissions = append(permissions, *permission)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.Internal(
			"Failed to iterate permissions",
			err,
		)
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

	return r.exists(
		ctx,
		query,
		name,
		excludeID,
		excludeID,
	)
}

func (r *permissionRepository) IsSlugExists(
	ctx context.Context,
	slug string,
	excludeID uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM permissions
			WHERE slug = ?
			  AND deleted_at IS NULL
			  AND (? = 0 OR id <> ?)
		)
	`

	return r.exists(
		ctx,
		query,
		slug,
		excludeID,
		excludeID,
	)
}

func (r *permissionRepository) IsPermissionExists(
	ctx context.Context,
	id uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM permissions
			WHERE id = ?
			  AND deleted_at IS NULL
		)
	`

	return r.exists(ctx, query, id)
}

func (r *permissionRepository) Create(
	ctx context.Context,
	p *entity.Permission,
) error {

	query := `
		INSERT INTO permissions (
			name,
			slug,
			group_name,
			description,
			created_by
		)
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		p.Name,
		p.Slug,
		p.GroupName,
		p.Description,
		p.CreatedBy,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to create permission",
			err,
		)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return apperror.Internal(
			"Failed to retrieve inserted permission ID",
			err,
		)
	}

	p.ID = uint64(id)

	return nil
}

func (r *permissionRepository) Update(
	ctx context.Context,
	p *entity.Permission,
) error {

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

	result, err := r.db.ExecContext(
		ctx,
		query,
		p.Name,
		p.GroupName,
		p.Description,
		p.UpdatedBy,
		p.ID,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to update permission",
			err,
		)
	}

	if _, err := result.RowsAffected(); err != nil {
		return apperror.Internal(
			"Failed to update permission",
			err,
		)
	}

	return nil
}

func (r *permissionRepository) Delete(
	ctx context.Context,
	id uint64,
	deletedBy uint64,
) error {

	query := `
		UPDATE permissions
		SET
			deleted_at = NOW(),
			deleted_by = ?
		WHERE id = ?
		  AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		deletedBy,
		id,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to delete permission",
			err,
		)
	}

	if _, err := result.RowsAffected(); err != nil {
		return apperror.Internal(
			"Failed to delete permission",
			err,
		)
	}

	return nil
}
