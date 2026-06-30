package repository

import (
	"context"
	"strings"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func (r *assignmentRepository) AssignUserRoles(
	ctx context.Context,
	userID uint64,
	roleIDs []uint64,
) error {

	// Nothing to assign.
	if len(roleIDs) == 0 {
		return nil
	}

	query := `
		INSERT INTO user_roles (
			user_id,
			role_id,
			created_at
		)
		VALUES
	`

	values := make([]string, 0, len(roleIDs))
	args := make([]any, 0, len(roleIDs)*3)

	now := time.Now()

	for _, roleID := range roleIDs {

		values = append(values, "(?, ?, ?)")

		args = append(
			args,
			userID,
			roleID,
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
		return apperror.Internal("Failed to assign user roles", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return apperror.Internal("Failed to retrieve affected rows", err)
	}

	return nil
}

func (r *assignmentRepository) ReplaceUserRoles(
	ctx context.Context,
	userID uint64,
	roleIDs []uint64,
) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return apperror.Internal("Failed to begin transaction", err)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	// Remove all existing user roles.
	_, err = tx.ExecContext(
		ctx,
		`
		DELETE
		FROM user_roles
		WHERE user_id = ?
		`,
		userID,
	)
	if err != nil {
		return apperror.Internal("Failed to delete user roles", err)
	}

	// Skip insert if there are no roles.
	if len(roleIDs) > 0 {

		query := `
			INSERT INTO user_roles (
				user_id,
				role_id,
				created_at
			)
			VALUES
		`

		values := make([]string, 0, len(roleIDs))
		args := make([]any, 0, len(roleIDs)*3)

		now := time.Now()

		for _, roleID := range roleIDs {

			values = append(values, "(?, ?, ?)")

			args = append(
				args,
				userID,
				roleID,
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
			return apperror.Internal("Failed to assign user roles", err)
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

func (r *assignmentRepository) FindUserRolesByUserID(
	ctx context.Context,
	userID uint64,
) ([]entity.UserRole, error) {

	// Retrieve all roles assigned to the user.
	query := `
		SELECT
			user_id,
			role_id,
			created_at,
			created_by
		FROM user_roles
		WHERE user_id = ?
		ORDER BY role_id ASC
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return nil, apperror.Internal("Failed to retrieve user roles", err)
	}
	defer rows.Close()

	roles := make([]entity.UserRole, 0)

	for rows.Next() {

		var role entity.UserRole

		if err := rows.Scan(
			&role.UserID,
			&role.RoleID,
			&role.CreatedAt,
			&role.CreatedBy,
		); err != nil {
			return nil, apperror.Internal("Failed to scan user role", err)
		}

		roles = append(
			roles,
			role,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.Internal("Failed to iterate user roles", err)
	}

	return roles, nil
}

func (r *assignmentRepository) DeleteUserRole(
	ctx context.Context,
	userID uint64,
	roleID uint64,
) error {

	// Remove a role from the user.
	result, err := r.db.ExecContext(
		ctx,
		`
		DELETE
		FROM user_roles
		WHERE user_id = ?
		AND role_id = ?
		`,
		userID,
		roleID,
	)
	if err != nil {
		return apperror.Internal("Failed to delete user role", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return apperror.Internal("Failed to retrieve affected rows", err)
	}

	return nil
}
