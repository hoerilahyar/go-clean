package bootstrap

import (
	"github.com/hoerilahyar/go-clean/internal/config"
	"github.com/hoerilahyar/go-clean/pkg/jwt"
)

type Services struct {
	JWT jwt.Service
}

func NewServices(cfg *config.Config) *Services {
	return &Services{
		JWT: jwt.New(
			cfg.JWTSecret,
			"go-clean",
			cfg.JWTAccessTokenTTL,
			cfg.JWTRefreshTokenTTL,
		),
	}
}
