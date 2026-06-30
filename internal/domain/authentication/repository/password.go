package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
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

	result, err := r.db.ExecContext(
		ctx,
		query,
		password,
		userID,
	)

	if err != nil {
		return apperror.Internal("Failed to update password", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return apperror.Internal("Failed to retrieve affected rows", err)
	}

	if rows == 0 {
		return nil
	}

	return nil
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

	if err != nil {
		return apperror.Internal("Failed to create password reset token", err)
	}

	return nil
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
			return nil, nil
		}

		return nil, apperror.Internal("Failed to retrieve password reset token", err)
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

	result, err := r.db.ExecContext(
		ctx,
		query,
		token,
	)

	if err != nil {
		return apperror.Internal("Failed to delete password reset token", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return apperror.Internal("Failed to retrieve affected rows", err)
	}

	return nil
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

	result, err := r.db.ExecContext(
		ctx,
		query,
		userID,
	)

	if err != nil {
		return apperror.Internal("Failed to delete password reset token", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return apperror.Internal("Failed to retrieve affected rows", err)
	}

	return nil
}
