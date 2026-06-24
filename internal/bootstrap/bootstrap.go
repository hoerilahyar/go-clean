package bootstrap

import (
	"github.com/hoerilahyar/go-clean/internal/config"
	permissionRepo "github.com/hoerilahyar/go-clean/internal/domain/permission/repository"
	permissionUsecase "github.com/hoerilahyar/go-clean/internal/domain/permission/usecase"
	roleRepo "github.com/hoerilahyar/go-clean/internal/domain/role/repository"
	roleUsecase "github.com/hoerilahyar/go-clean/internal/domain/role/usecase"
	userRepo "github.com/hoerilahyar/go-clean/internal/domain/user/repository"
	userUsecase "github.com/hoerilahyar/go-clean/internal/domain/user/usecase"
	"github.com/hoerilahyar/go-clean/internal/handler/http"
	"github.com/hoerilahyar/go-clean/internal/infrastructure/database"
)

func NewApplication() *Application {
	cfg := config.Load()

	db := database.NewMySQLConnection(cfg)

	userRepo := userRepo.NewUserRepository(db)
	roleRepo := roleRepo.NewRoleRepository(db)
	permissionRepo := permissionRepo.NewPermissionRepository(db)

	userUC := userUsecase.NewUserUsecase(userRepo)
	roleUC := roleUsecase.NewRoleUsecase(roleRepo)
	permissionUC := permissionUsecase.NewPermissionUsecase(permissionRepo)

	return &Application{
		Config: cfg,

		UserHandler:       http.NewUserHandler(userUC),
		RoleHandler:       http.NewRoleHandler(roleUC),
		PermissionHandler: http.NewPermissionHandler(permissionUC),
	}
}
