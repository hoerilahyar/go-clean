package bootstrap

import (
	"github.com/hoerilahyar/go-clean/internal/config"
	permissionHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/handler"
	permissionRepo "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/repository"
	permissionUsecase "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/usecase"
	roleHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/handler"
	roleRepo "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/repository"
	roleUsecase "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/usecase"
	userHandler "github.com/hoerilahyar/go-clean/internal/domain/user/handler"
	userRepo "github.com/hoerilahyar/go-clean/internal/domain/user/repository"
	userUsecase "github.com/hoerilahyar/go-clean/internal/domain/user/usecase"
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

		UserHandler:       userHandler.NewUserHandler(userUC),
		RoleHandler:       roleHandler.NewRoleHandler(roleUC),
		PermissionHandler: permissionHandler.NewPermissionHandler(permissionUC),
	}
}
