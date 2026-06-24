package bootstrap

import (
	"database/sql"

	"github.com/hoerilahyar/go-clean/internal/config"
	"github.com/hoerilahyar/go-clean/internal/handler/http"
)

type Application struct {
	Config *config.Config
	DB     *sql.DB

	UserHandler       *http.UserHandler
	RoleHandler       *http.RoleHandler
	PermissionHandler *http.PermissionHandler
}
