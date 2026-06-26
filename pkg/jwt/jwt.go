package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateAccessToken(userID uint64, username string) (string, error)
	GenerateRefreshToken() (string, time.Time, error)

	ParseAccessToken(token string) (*Claims, error)
	ValidateAccessToken(token string) (*Claims, error)

	AccessTokenTTL() int64
	RefreshTokenTTL() time.Duration
}

type service struct {
	secretKey       []byte
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	issuer          string
}

func New(
	secretKey string,
	issuer string,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) Service {

	return &service{
		secretKey:       []byte(secretKey),
		issuer:          issuer,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (s *service) AccessTokenTTL() int64 {
	return int64(s.accessTokenTTL.Seconds())
}

func (s *service) RefreshTokenTTL() time.Duration {
	return s.refreshTokenTTL
}

var _ jwt.Claims
