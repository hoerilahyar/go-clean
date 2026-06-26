package bootstrap

import (
	"github.com/hoerilahyar/go-clean/internal/config"

	assignmentHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/handler"
	assignmentRepo "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/repository"
	assignmentUsecase "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/usecase"

	permissionHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/handler"
	permissionRepo "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/repository"
	permissionUsecase "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/usecase"

	roleHandler "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/handler"
	roleRepo "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/repository"
	roleUsecase "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/usecase"

	userHandler "github.com/hoerilahyar/go-clean/internal/domain/user/handler"
	userRepo "github.com/hoerilahyar/go-clean/internal/domain/user/repository"
	userUsecase "github.com/hoerilahyar/go-clean/internal/domain/user/usecase"

	authNHandler "github.com/hoerilahyar/go-clean/internal/domain/authentication/handler"
	authNRepo "github.com/hoerilahyar/go-clean/internal/domain/authentication/repository"
	authNUsecase "github.com/hoerilahyar/go-clean/internal/domain/authentication/usecase"

	"github.com/hoerilahyar/go-clean/internal/infrastructure/database"
)

func NewApplication() *Application {
	cfg := config.Load()

	db := database.NewMySQLConnection(cfg)

	services := NewServices(cfg)

	// Repository
	userRepository := userRepo.NewUserRepository(db)
	roleRepository := roleRepo.NewRoleRepository(db)
	permissionRepository := permissionRepo.NewPermissionRepository(db)
	assignmentRepository := assignmentRepo.NewAssignmentRepository(db)
	authenticationRepo := authNRepo.NewAuthenticationRepository(db)

	// Usecase
	userUC := userUsecase.NewUserUsecase(userRepository)
	roleUC := roleUsecase.NewRoleUsecase(roleRepository)
	permissionUC := permissionUsecase.NewPermissionUsecase(permissionRepository)
	assignmentUC := assignmentUsecase.NewAssignmentUsecase(assignmentRepository)
	authenticationUC := authNUsecase.NewAuthenticationUsecase(authenticationRepo, services.JWT)

	// Handler
	user := userHandler.NewUserHandler(userUC)

	authorize := &AuthorizeHandler{
		Role:       roleHandler.NewRoleHandler(roleUC),
		Permission: permissionHandler.NewPermissionHandler(permissionUC),
		Assignment: assignmentHandler.NewAssignmentHandler(assignmentUC),
	}

	authentication := authNHandler.NewAuthenticationHandler(authenticationUC)

	return &Application{
		Config: cfg,
		DB:     db,

		Services: services,

		User: user,

		Authorize:      authorize,
		Authentication: authentication,
	}
}
