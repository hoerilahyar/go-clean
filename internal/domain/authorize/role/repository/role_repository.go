package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

const roleColumns = `
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
`

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(
	db *sql.DB,
) RoleRepository {

	return &roleRepository{
		db: db,
	}
}

// scanRole scans a role from sql.Row or sql.Rows.
func scanRole(scanner interface {
	Scan(dest ...any) error
}) (*entity.Role, error) {

	role := &entity.Role{}

	err := scanner.Scan(
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

// findOne executes a query that returns a single role.
func (r *roleRepository) findOne(
	ctx context.Context,
	query string,
	args ...any,
) (*entity.Role, error) {

	row := r.db.QueryRowContext(
		ctx,
		query,
		args...,
	)

	role, err := scanRole(row)
	if err != nil {
		return nil, apperror.Internal(
			"Failed to retrieve role",
			err,
		)
	}

	return role, nil
}

// exists executes an EXISTS query.
func (r *roleRepository) exists(
	ctx context.Context,
	query string,
	args ...any,
) (bool, error) {

	var exists bool

	err := r.db.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(&exists)

	if err != nil {
		return false, apperror.Internal(
			"Failed to check role existence",
			err,
		)
	}

	return exists, nil
}

// buildPlaceholders builds placeholders for IN queries.
func buildPlaceholders(
	count int,
) string {

	placeholders := make(
		[]string,
		count,
	)

	for i := range placeholders {
		placeholders[i] = "?"
	}

	return strings.Join(
		placeholders,
		",",
	)
}
func (r *roleRepository) FindAll(
	ctx context.Context,
	req request.GetRolesRequest,
) ([]entity.Role, error) {

	query := `
		SELECT
			` + roleColumns + `
		FROM roles
		WHERE deleted_at IS NULL
	`

	args := make([]any, 0)

	if req.ID != nil {
		query += " AND id = ?"
		args = append(args, *req.ID)
	}

	if req.Name != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+req.Name+"%")
	}

	if req.Slug != "" {
		query += " AND slug LIKE ?"
		args = append(args, "%"+req.Slug+"%")
	}

	if req.IsActive != nil {
		query += " AND is_active = ?"
		args = append(args, *req.IsActive)
	}

	query += " ORDER BY name ASC"

	rows, err := r.db.QueryContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, apperror.Internal(
			"Failed to retrieve roles",
			err,
		)
	}

	defer rows.Close()

	roles := make([]entity.Role, 0)

	for rows.Next() {

		role, err := scanRole(rows)
		if err != nil {
			return nil, apperror.Internal(
				"Failed to scan role",
				err,
			)
		}

		roles = append(
			roles,
			*role,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.Internal(
			"Failed to iterate roles",
			err,
		)
	}

	return roles, nil
}

func (r *roleRepository) FindByID(
	ctx context.Context,
	id uint64,
) (*entity.Role, error) {

	query := `
		SELECT
			` + roleColumns + `
		FROM roles
		WHERE id = ?
		  AND deleted_at IS NULL
	`

	return r.findOne(
		ctx,
		query,
		id,
	)
}

func (r *roleRepository) FindBySlug(
	ctx context.Context,
	slug string,
) (*entity.Role, error) {

	query := `
		SELECT
			` + roleColumns + `
		FROM roles
		WHERE slug = ?
		  AND deleted_at IS NULL
	`

	return r.findOne(
		ctx,
		query,
		slug,
	)
}

func (r *roleRepository) FindByIDs(
	ctx context.Context,
	ids []uint64,
) ([]entity.Role, error) {

	if len(ids) == 0 {
		return []entity.Role{}, nil
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
			description,
			created_at,
			updated_at
		FROM roles
		WHERE id IN (` + buildPlaceholders(len(ids)) + `)
		  AND deleted_at IS NULL
		ORDER BY id ASC
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, apperror.Internal(
			"Failed to retrieve roles",
			err,
		)
	}

	defer rows.Close()

	roles := make([]entity.Role, 0, len(ids))

	for rows.Next() {

		var role entity.Role

		err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.Slug,
			&role.Description,
			&role.CreatedAt,
			&role.UpdatedAt,
		)
		if err != nil {
			return nil, apperror.Internal(
				"Failed to scan role",
				err,
			)
		}

		roles = append(
			roles,
			role,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.Internal(
			"Failed to iterate roles",
			err,
		)
	}

	return roles, nil
}

func (r *roleRepository) IsSlugExists(
	ctx context.Context,
	slug string,
	excludeID uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM roles
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

func (r *roleRepository) IsNameExists(
	ctx context.Context,
	name string,
	excludeID uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM roles
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

func (r *roleRepository) IsRoleExists(
	ctx context.Context,
	id uint64,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM roles
			WHERE id = ?
			  AND deleted_at IS NULL
		)
	`

	return r.exists(
		ctx,
		query,
		id,
	)
}

func (r *roleRepository) Create(
	ctx context.Context,
	role *entity.Role,
) error {

	query := `
		INSERT INTO roles (
			name,
			slug,
			description,
			is_active,
			created_by
		)
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		role.Name,
		role.Slug,
		role.Description,
		role.IsActive,
		role.CreatedBy,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to create role",
			err,
		)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return apperror.Internal(
			"Failed to retrieve inserted role ID",
			err,
		)
	}

	role.ID = uint64(id)

	return nil
}

func (r *roleRepository) Update(
	ctx context.Context,
	role *entity.Role,
) error {

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

	result, err := r.db.ExecContext(
		ctx,
		query,
		role.Name,
		role.Slug,
		role.Description,
		role.IsActive,
		role.UpdatedBy,
		role.ID,
	)
	if err != nil {
		return apperror.Internal(
			"Failed to update role",
			err,
		)
	}

	if _, err := result.RowsAffected(); err != nil {
		return apperror.Internal(
			"Failed to retrieve affected rows",
			err,
		)
	}

	return nil
}

func (r *roleRepository) Delete(
	ctx context.Context,
	id uint64,
	deletedBy uint64,
) error {

	query := `
		UPDATE roles
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
			"Failed to delete role",
			err,
		)
	}

	if _, err := result.RowsAffected(); err != nil {
		return apperror.Internal(
			"Failed to retrieve affected rows",
			err,
		)
	}

	return nil
}
