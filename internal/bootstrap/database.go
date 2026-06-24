package bootstrap

import (
	"database/sql"

	"github.com/hoerilahyar/go-clean/internal/config"
	"github.com/hoerilahyar/go-clean/internal/infrastructure/database"
)

func NewDatabase(cfg *config.Config) *sql.DB {
	return database.NewMySQLConnection(cfg)
}
