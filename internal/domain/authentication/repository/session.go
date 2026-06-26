package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/entity"
)

func (r *authenticationRepository) CreateSession(
	ctx context.Context,
	session entity.Session,
) error {

	query := `
		INSERT INTO user_sessions (
			user_id,
			refresh_token,
			user_agent,
			ip_address,
			expired_at,
			created_at
		)
		VALUES (?, ?, ?, ?, ?, NOW())
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		session.UserID,
		session.RefreshToken,
		session.UserAgent,
		session.IPAddress,
		session.ExpiredAt,
	)

	return err
}

func (r *authenticationRepository) FindSessionByRefreshToken(
	ctx context.Context,
	refreshToken string,
) (*entity.Session, error) {

	query := `
		SELECT
			id,
			user_id,
			refresh_token,
			user_agent,
			ip_address,
			expired_at,
			created_at
		FROM user_sessions
		WHERE refresh_token = ?
		AND deleted_at IS NULL
		LIMIT 1
	`

	var session entity.Session

	err := r.db.QueryRowContext(
		ctx,
		query,
		refreshToken,
	).Scan(
		&session.ID,
		&session.UserID,
		&session.RefreshToken,
		&session.UserAgent,
		&session.IPAddress,
		&session.ExpiredAt,
		&session.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, err
	}

	return &session, nil
}

func (r *authenticationRepository) UpdateSession(
	ctx context.Context,
	session entity.Session,
) error {

	query := `
		UPDATE user_sessions
		SET
			refresh_token = ?,
			expired_at = ?,
			updated_at = NOW()
		WHERE id = ?
		AND deleted_at IS NULL
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		session.RefreshToken,
		session.ExpiredAt,
		session.ID,
	)

	return err
}

func (r *authenticationRepository) DeleteSessionByRefreshToken(
	ctx context.Context,
	refreshToken string,
) error {

	query := `
		DELETE
		FROM user_sessions
		WHERE refresh_token = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		refreshToken,
	)

	return err
}

func (r *authenticationRepository) DeleteSessionsByUserID(
	ctx context.Context,
	userID uint64,
) error {

	query := `
		DELETE
		FROM user_sessions
		WHERE user_id = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		userID,
	)

	return err
}
