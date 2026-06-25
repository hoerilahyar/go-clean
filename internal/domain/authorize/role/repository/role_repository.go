package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authorize/role/entity"
)

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) FindAll(ctx context.Context, req request.GetRolesRequest) ([]entity.Role, error) {
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

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := make([]entity.Role, 0)

	for rows.Next() {
		var role entity.Role

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

func (r *roleRepository) FindByIDs(
	ctx context.Context,
	ids []uint64,
) ([]*entity.Role, error) {

	if len(ids) == 0 {
		return []*entity.Role{}, nil
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
			description,
			created_at,
			updated_at
		FROM roles
		WHERE id IN (%s)
		AND deleted_at IS NULL
		ORDER BY id ASC
	`, strings.Join(placeholders, ","))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*entity.Role

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
			return nil, err
		}

		roles = append(roles, &role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
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

func (r *roleRepository) IsSlugExists(ctx context.Context, slug string, excludeID uint64) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM roles
			WHERE slug = ?
				AND deleted_at IS NULL
				AND (? = 0 OR id <> ?)
		)
	`

	var exists bool

	err := r.db.QueryRowContext(ctx, query, slug, excludeID, excludeID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *roleRepository) IsNameExists(ctx context.Context, name string, excludeID uint64) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM roles
			WHERE name = ?
				AND deleted_at IS NULL
				AND (? = 0 OR id <> ?)
		)
	`

	var exists bool

	err := r.db.QueryRowContext(ctx, query, name, excludeID, excludeID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *roleRepository) IsRoleExists(ctx context.Context, id uint64) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM roles
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

func (r *roleRepository) Create(ctx context.Context, role *entity.Role) error {
	fmt.Println("Creating role:", role)
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
