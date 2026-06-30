package repository

import (
	"context"
	"strings"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func (r *assignmentRepository) AssignUserPermissions(
	ctx context.Context,
	userID uint64,
	permissionIDs []uint64,
) error {

	// Nothing to assign.
	if len(permissionIDs) == 0 {
		return nil
	}

	query := `
		INSERT INTO user_permissions (
			user_id,
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
			userID,
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
		return apperror.Internal("Failed to assign user permissions", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return apperror.Internal("Failed to retrieve affected rows", err)
	}

	return nil
}

func (r *assignmentRepository) ReplaceUserPermissions(
	ctx context.Context,
	userID uint64,
	permissionIDs []uint64,
) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return apperror.Internal("Failed to begin transaction", err)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	// Remove all existing user permissions.
	_, err = tx.ExecContext(
		ctx,
		`
		DELETE
		FROM user_permissions
		WHERE user_id = ?
		`,
		userID,
	)
	if err != nil {
		return apperror.Internal("Failed to delete user permissions", err)
	}

	// Skip insert if there are no permissions.
	if len(permissionIDs) > 0 {

		query := `
			INSERT INTO user_permissions (
				user_id,
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
				userID,
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
			return apperror.Internal("Failed to assign user permissions", err)
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

func (r *assignmentRepository) FindUserPermissionsByUserID(
	ctx context.Context,
	userID uint64,
) ([]entity.UserPermission, error) {

	// Retrieve all permissions assigned to the user.
	query := `
		SELECT
			user_id,
			permission_id,
			created_at,
			created_by
		FROM user_permissions
		WHERE user_id = ?
		ORDER BY permission_id ASC
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return nil, apperror.Internal("Failed to retrieve user permissions", err)
	}
	defer rows.Close()

	permissions := make([]entity.UserPermission, 0)

	for rows.Next() {

		var permission entity.UserPermission

		if err := rows.Scan(
			&permission.UserID,
			&permission.PermissionID,
			&permission.CreatedAt,
			&permission.CreatedBy,
		); err != nil {
			return nil, apperror.Internal("Failed to scan user permission", err)
		}

		permissions = append(
			permissions,
			permission,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.Internal("Failed to iterate user permissions", err)
	}

	return permissions, nil
}

func (r *assignmentRepository) DeleteUserPermission(
	ctx context.Context,
	userID uint64,
	permissionID uint64,
) error {

	// Remove a permission from the user.
	result, err := r.db.ExecContext(
		ctx,
		`
		DELETE
		FROM user_permissions
		WHERE user_id = ?
		AND permission_id = ?
		`,
		userID,
		permissionID,
	)
	if err != nil {
		return apperror.Internal("Failed to delete user permission", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return apperror.Internal("Failed to retrieve affected rows", err)
	}

	return nil
}
