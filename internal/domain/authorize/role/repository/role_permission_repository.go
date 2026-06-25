package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	permissionEntity "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/entity"
)

type rolePermissionRepository struct {
	db *sql.DB
}

func NewRolePermissionRepository(db *sql.DB) RolePermissionRepository {
	return &rolePermissionRepository{
		db: db,
	}
}

func (r *rolePermissionRepository) GetPermissions(
	ctx context.Context,
	roleID uint64,
) ([]*permissionEntity.Permission, error) {

	query := `
	SELECT
		p.id,
		p.name,
		p.slug,
		p.group_name,
		p.created_at,
		p.updated_at
	FROM role_permissions rp
	JOIN permissions p ON p.id = rp.permission_id
	WHERE rp.role_id = ?
	AND rp.deleted_at IS NULL
	AND p.deleted_at IS NULL
	ORDER BY p.id ASC
	`

	rows, err := r.db.QueryContext(ctx, query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*permissionEntity.Permission

	for rows.Next() {
		var p permissionEntity.Permission

		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Slug,
			&p.GroupName,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}

		permissions = append(permissions, &p)
	}

	return permissions, rows.Err()
}

func (r *rolePermissionRepository) Assign(
	ctx context.Context,
	roleID uint64,
	permissionIDs []uint64,
) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
	INSERT IGNORE INTO role_permissions
	(role_id, permission_id)
	VALUES (?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, permissionID := range permissionIDs {
		if _, err := stmt.ExecContext(
			ctx,
			roleID,
			permissionID,
		); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *rolePermissionRepository) Remove(
	ctx context.Context,
	roleID uint64,
	permissionID uint64,
) error {

	_, err := r.db.ExecContext(
		ctx,
		`DELETE FROM role_permissions
		 WHERE role_id = ?
		 AND permission_id = ?`,
		roleID,
		permissionID,
	)

	return err
}

func (r *rolePermissionRepository) Sync(
	ctx context.Context,
	roleID uint64,
	permissionIDs []uint64,
) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if len(permissionIDs) == 0 {

		_, err = tx.ExecContext(
			ctx,
			`DELETE FROM role_permissions
			 WHERE role_id = ?`,
			roleID,
		)

		if err != nil {
			return err
		}

		return tx.Commit()
	}

	placeholders := make([]string, len(permissionIDs))
	args := make([]any, 0, len(permissionIDs)+1)

	args = append(args, roleID)

	for i, id := range permissionIDs {
		placeholders[i] = "?"
		args = append(args, id)
	}

	query := fmt.Sprintf(
		`DELETE FROM role_permissions
		 WHERE role_id = ?
		 AND permission_id NOT IN (%s)`,
		strings.Join(placeholders, ","),
	)

	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, `
	INSERT IGNORE INTO role_permissions
	(role_id, permission_id)
	VALUES (?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, permissionID := range permissionIDs {
		if _, err := stmt.ExecContext(
			ctx,
			roleID,
			permissionID,
		); err != nil {
			return err
		}
	}

	return tx.Commit()
}
