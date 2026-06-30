package repository

import (
	"context"
	"database/sql"
	"errors"

	userEntity "github.com/hoerilahyar/go-clean/internal/domain/user/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func (r *authenticationRepository) FindUserByIdentity(
	ctx context.Context,
	identity string,
) (*userEntity.User, error) {

	query := `
		SELECT
			id,
			full_name,
			username,
			email,
			password,
			status
		FROM users
		WHERE (
			username = ?
			OR email = ?
		)
		AND deleted_at IS NULL
		LIMIT 1
	`

	var user userEntity.User

	err := r.db.QueryRowContext(
		ctx,
		query,
		identity,
		identity,
	).Scan(
		&user.ID,
		&user.FullName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Status,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, apperror.Internal("Failed to retrieve user", err)
	}

	return &user, nil
}

func (r *authenticationRepository) FindUserByID(
	ctx context.Context,
	userID uint64,
) (*userEntity.User, error) {

	query := `
		SELECT
			id,
			full_name,
			username,
			email,
			password,
			status
		FROM users
		WHERE id = ?
		AND deleted_at IS NULL
		LIMIT 1
	`

	var user userEntity.User

	err := r.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.FullName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Status,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, apperror.Internal("Failed to retrieve user", err)
	}

	return &user, nil
}
