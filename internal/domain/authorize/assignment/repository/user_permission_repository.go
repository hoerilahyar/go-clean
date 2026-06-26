package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
)

func (r *assignmentRepository) AssignUserPermissions(
	ctx context.Context,
	userID uint64,
	permissionIDs []uint64,
) error {

	if len(permissionIDs) == 0 {
		return nil
	}

	query := `
		INSERT INTO user_permissions
		(
			user_id,
			permission_id,
			created_at
		)
		VALUES
	`

	args := make([]any, 0)
	values := make([]string, 0)

	now := time.Now()

	for _, permissionID := range permissionIDs {

		values = append(values, "(?, ?, ?)")

		args = append(args,
			userID,
			permissionID,
			now,
		)
	}

	query += strings.Join(values, ",")

	_, err := r.db.ExecContext(ctx, query, args...)

	return err
}

func (r *assignmentRepository) ReplaceUserPermissions(
	ctx context.Context,
	userID uint64,
	permissionIDs []uint64,
) error {

	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(
		ctx,
		`
		DELETE FROM user_permissions
		WHERE user_id = ?
		`,
		userID,
	)

	if err != nil {
		return err
	}

	if len(permissionIDs) > 0 {

		query := `
			INSERT INTO user_permissions
			(
				user_id,
				permission_id,
				created_at
			)
			VALUES
		`

		args := make([]any, 0)
		values := make([]string, 0)

		now := time.Now()

		for _, permissionID := range permissionIDs {

			values = append(values, "(?, ?, ?)")

			args = append(args,
				userID,
				permissionID,
				now,
			)
		}

		query += strings.Join(values, ",")

		_, err = tx.ExecContext(ctx, query, args...)

		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *assignmentRepository) FindUserPermissionsByUserID(
	ctx context.Context,
	userID uint64,
) ([]entity.UserPermission, error) {

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

	rows, err := r.db.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var permissions []entity.UserPermission

	for rows.Next() {

		var permission entity.UserPermission

		err = rows.Scan(
			&permission.UserID,
			&permission.PermissionID,
			&permission.CreatedAt,
			&permission.CreatedBy,
		)

		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *assignmentRepository) DeleteUserPermission(
	ctx context.Context,
	userID uint64,
	permissionID uint64,
) error {

	result, err := r.db.ExecContext(
		ctx,
		`
		DELETE FROM user_permissions
		WHERE user_id = ?
		AND permission_id = ?
		`,
		userID,
		permissionID,
	)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("user permission not found")
	}

	return nil
}
