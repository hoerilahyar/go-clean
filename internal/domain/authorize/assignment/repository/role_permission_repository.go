package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
)

func (r *assignmentRepository) AssignRolePermissions(
	ctx context.Context,
	roleID uint64,
	permissionIDs []uint64,
) error {

	if len(permissionIDs) == 0 {
		return nil
	}

	query := `
		INSERT INTO role_permissions
		(
			role_id,
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
			roleID,
			permissionID,
			now,
		)
	}

	query += strings.Join(values, ",")

	_, err := r.db.ExecContext(ctx, query, args...)

	return err
}

func (r *assignmentRepository) ReplaceRolePermissions(
	ctx context.Context,
	roleID uint64,
	permissionIDs []uint64,
) error {

	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(
		ctx,
		`DELETE FROM role_permissions WHERE role_id = ?`,
		roleID,
	)

	if err != nil {
		return err
	}

	if len(permissionIDs) > 0 {

		query := `
			INSERT INTO role_permissions
			(
				role_id,
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
				roleID,
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

func (r *assignmentRepository) FindRolePermissionsByRoleID(
	ctx context.Context,
	roleID uint64,
) ([]entity.RolePermission, error) {

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

	rows, err := r.db.QueryContext(ctx, query, roleID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var permissions []entity.RolePermission

	for rows.Next() {

		var permission entity.RolePermission

		err = rows.Scan(
			&permission.RoleID,
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

func (r *assignmentRepository) DeleteRolePermission(
	ctx context.Context,
	roleID uint64,
	permissionID uint64,
) error {

	result, err := r.db.ExecContext(
		ctx,
		`
		DELETE FROM role_permissions
		WHERE role_id = ?
		AND permission_id = ?
		`,
		roleID,
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
		return fmt.Errorf("role permission not found")
	}

	return nil
}
