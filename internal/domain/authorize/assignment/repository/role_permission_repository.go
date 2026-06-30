package repository

import (
	"context"
	"strings"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func (r *assignmentRepository) AssignRolePermissions(
	ctx context.Context,
	roleID uint64,
	permissionIDs []uint64,
) error {

	// Nothing to assign.
	if len(permissionIDs) == 0 {
		return nil
	}

	query := `
		INSERT INTO role_permissions (
			role_id,
			permission_id,
			created_at
		)
		VALUES
	`

	values := make([]string, 0, len(permissionIDs))
	args := make([]any, 0, len(permissionIDs)*3)

	now := time.Now()

	for _, permissionID := range permissionIDs {

		values = append(values, "(?, ?, ?)")

		args = append(
			args,
			roleID,
			permissionID,
			now,
		)
	}

	query += strings.Join(values, ",")

	result, err := r.db.ExecContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return apperror.Internal("Failed to assign role permissions", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return apperror.Internal("Failed to retrieve affected rows", err)
	}

	return nil
}

func (r *assignmentRepository) ReplaceRolePermissions(
	ctx context.Context,
	roleID uint64,
	permissionIDs []uint64,
) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return apperror.Internal("Failed to begin transaction", err)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	// Remove all existing role permissions.
	_, err = tx.ExecContext(
		ctx,
		`
		DELETE
		FROM role_permissions
		WHERE role_id = ?
		`,
		roleID,
	)
	if err != nil {
		return apperror.Internal("Failed to delete role permissions", err)
	}

	// Skip insert if there are no permissions.
	if len(permissionIDs) > 0 {

		query := `
			INSERT INTO role_permissions (
				role_id,
				permission_id,
				created_at
			)
			VALUES
		`

		values := make([]string, 0, len(permissionIDs))
		args := make([]any, 0, len(permissionIDs)*3)

		now := time.Now()

		for _, permissionID := range permissionIDs {

			values = append(values, "(?, ?, ?)")

			args = append(
				args,
				roleID,
				permissionID,
				now,
			)
		}

		query += strings.Join(values, ",")

		result, err := tx.ExecContext(
			ctx,
			query,
			args...,
		)
		if err != nil {
			return apperror.Internal("Failed to assign role permissions", err)
		}

		_, err = result.RowsAffected()
		if err != nil {
			return apperror.Internal("Failed to retrieve affected rows", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return apperror.Internal("Failed to commit transaction", err)
	}

	return nil
}

func (r *assignmentRepository) FindRolePermissionsByRoleID(
	ctx context.Context,
	roleID uint64,
) ([]entity.RolePermission, error) {

	// Retrieve all permissions assigned to the role.
	query := `
		SELECT
			role_id,
			permission_id,
			created_at,
			created_by
		FROM role_permissions
		WHERE role_id = ?
		ORDER BY permission_id ASC
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		roleID,
	)
	if err != nil {
		return nil, apperror.Internal("Failed to retrieve role permissions", err)
	}
	defer rows.Close()

	permissions := make([]entity.RolePermission, 0)

	for rows.Next() {

		var permission entity.RolePermission

		if err := rows.Scan(
			&permission.RoleID,
			&permission.PermissionID,
			&permission.CreatedAt,
			&permission.CreatedBy,
		); err != nil {
			return nil, apperror.Internal("Failed to scan role permission", err)
		}

		permissions = append(
			permissions,
			permission,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.Internal("Failed to iterate role permissions", err)
	}

	return permissions, nil
}

func (r *assignmentRepository) DeleteRolePermission(
	ctx context.Context,
	roleID uint64,
	permissionID uint64,
) error {

	// Remove a permission from the role.
	result, err := r.db.ExecContext(
		ctx,
		`
		DELETE
		FROM role_permissions
		WHERE role_id = ?
		AND permission_id = ?
		`,
		roleID,
		permissionID,
	)
	if err != nil {
		return apperror.Internal("Failed to delete role permission", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return apperror.Internal("Failed to retrieve affected rows", err)
	}

	return nil
}
