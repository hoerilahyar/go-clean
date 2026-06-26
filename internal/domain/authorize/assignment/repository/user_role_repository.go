package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/entity"
)

func (r *assignmentRepository) AssignUserRoles(
	ctx context.Context,
	userID uint64,
	roleIDs []uint64,
) error {

	if len(roleIDs) == 0 {
		return nil
	}

	query := `
		INSERT INTO user_roles
		(
			user_id,
			role_id,
			created_at
		)
		VALUES
	`

	args := make([]any, 0)
	values := make([]string, 0)

	now := time.Now()

	for _, roleID := range roleIDs {

		values = append(values, "(?, ?, ?)")

		args = append(args,
			userID,
			roleID,
			now,
		)
	}

	query += strings.Join(values, ",")

	_, err := r.db.ExecContext(ctx, query, args...)

	return err
}

func (r *assignmentRepository) ReplaceUserRoles(
	ctx context.Context,
	userID uint64,
	roleIDs []uint64,
) error {

	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(
		ctx,
		`
		DELETE FROM user_roles
		WHERE user_id = ?
		`,
		userID,
	)

	if err != nil {
		return err
	}

	if len(roleIDs) > 0 {

		query := `
			INSERT INTO user_roles
			(
				user_id,
				role_id,
				created_at
			)
			VALUES
		`

		args := make([]any, 0)
		values := make([]string, 0)

		now := time.Now()

		for _, roleID := range roleIDs {

			values = append(values, "(?, ?, ?)")

			args = append(args,
				userID,
				roleID,
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

func (r *assignmentRepository) FindUserRolesByUserID(
	ctx context.Context,
	userID uint64,
) ([]entity.UserRole, error) {

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

	rows, err := r.db.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var roles []entity.UserRole

	for rows.Next() {

		var role entity.UserRole

		err = rows.Scan(
			&role.UserID,
			&role.RoleID,
			&role.CreatedAt,
			&role.CreatedBy,
		)

		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *assignmentRepository) DeleteUserRole(
	ctx context.Context,
	userID uint64,
	roleID uint64,
) error {

	result, err := r.db.ExecContext(
		ctx,
		`
		DELETE FROM user_roles
		WHERE user_id = ?
		AND role_id = ?
		`,
		userID,
		roleID,
	)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("user role not found")
	}

	return nil
}
