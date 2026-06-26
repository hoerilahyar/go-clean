package repository

import (
	"context"
	"database/sql"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/entity"
	userEntity "github.com/hoerilahyar/go-clean/internal/domain/user/entity"
)

type AuthenticationRepository interface {
	// Login
	FindUserByIdentity(ctx context.Context, identity string) (*userEntity.User, error)
	FindUserByID(ctx context.Context, userID uint64) (*userEntity.User, error)

	// Session
	CreateSession(ctx context.Context, session entity.Session) error
	FindSessionByRefreshToken(ctx context.Context, refreshToken string) (*entity.Session, error)
	DeleteSessionByRefreshToken(ctx context.Context, refreshToken string) error
	DeleteSessionsByUserID(ctx context.Context, userID uint64) error
	UpdateSession(ctx context.Context, session entity.Session) error

	// Password
	UpdatePassword(ctx context.Context, userID uint64, password string) error

	// Password Reset
	CreatePasswordResetToken(ctx context.Context, reset entity.PasswordReset) error
	FindPasswordResetToken(ctx context.Context, token string) (*entity.PasswordReset, error)
	DeletePasswordResetToken(ctx context.Context, token string) error
	DeletePasswordResetTokenByUserID(ctx context.Context, userID uint64) error
}

type authenticationRepository struct {
	db *sql.DB
}

func NewAuthenticationRepository(
	db *sql.DB,
) AuthenticationRepository {
	return &authenticationRepository{
		db: db,
	}
}
