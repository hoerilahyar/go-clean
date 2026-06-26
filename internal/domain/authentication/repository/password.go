package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/entity"
)

func (r *authenticationRepository) UpdatePassword(
	ctx context.Context,
	userID uint64,
	password string,
) error {

	query := `
		UPDATE users
		SET
			password = ?,
			updated_at = NOW()
		WHERE id = ?
		AND deleted_at IS NULL
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		password,
		userID,
	)

	return err
}

func (r *authenticationRepository) CreatePasswordResetToken(
	ctx context.Context,
	reset entity.PasswordReset,
) error {

	query := `
		INSERT INTO password_reset_tokens (
			user_id,
			token,
			expired_at,
			created_at
		)
		VALUES (?, ?, ?, NOW())
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		reset.UserID,
		reset.Token,
		reset.ExpiredAt,
	)

	return err
}

func (r *authenticationRepository) FindPasswordResetToken(
	ctx context.Context,
	token string,
) (*entity.PasswordReset, error) {

	query := `
		SELECT
			id,
			user_id,
			token,
			expired_at,
			created_at
		FROM password_reset_tokens
		WHERE token = ?
		LIMIT 1
	`

	var reset entity.PasswordReset

	err := r.db.QueryRowContext(
		ctx,
		query,
		token,
	).Scan(
		&reset.ID,
		&reset.UserID,
		&reset.Token,
		&reset.ExpiredAt,
		&reset.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, err
	}

	return &reset, nil
}

func (r *authenticationRepository) DeletePasswordResetToken(
	ctx context.Context,
	token string,
) error {

	query := `
		DELETE
		FROM password_reset_tokens
		WHERE token = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		token,
	)

	return err
}

func (r *authenticationRepository) DeletePasswordResetTokenByUserID(
	ctx context.Context,
	userID uint64,
) error {

	query := `
		DELETE
		FROM password_reset_tokens
		WHERE user_id = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		userID,
	)

	return err
}
