package bootstrap

import (
	"time"

	"github.com/hoerilahyar/go-clean/internal/config"
	authNUsecase "github.com/hoerilahyar/go-clean/internal/domain/authentication/usecase"
	assignmentUsecase "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/usecase"
	permissionUsecase "github.com/hoerilahyar/go-clean/internal/domain/authorize/permission/usecase"
	roleUsecase "github.com/hoerilahyar/go-clean/internal/domain/authorize/role/usecase"
	userUsecase "github.com/hoerilahyar/go-clean/internal/domain/user/usecase"
	"github.com/hoerilahyar/go-clean/pkg/jwt"
)

type Usecases struct {
	User userUsecase.UserUsecase

	Role       roleUsecase.RoleUsecase
	Permission permissionUsecase.PermissionUsecase
	Assignment assignmentUsecase.AssignmentUsecase
	AuthN      authNUsecase.AuthenticationUsecase
}

func NewUsecases(
	repo *Repositories,
	cfg *config.Config,
) *Usecases {

	jwtService := jwt.New(
		cfg.JWTSecret,
		cfg.JWTIssuer,
		time.Duration(cfg.JWTAccessTokenTTL)*time.Second,
		time.Duration(cfg.JWTRefreshTokenTTL)*time.Second,
	)

	return &Usecases{
		User: userUsecase.NewUserUsecase(repo.User),

		Role:       roleUsecase.NewRoleUsecase(repo.Role),
		Permission: permissionUsecase.NewPermissionUsecase(repo.Permission),

		Assignment: assignmentUsecase.NewAssignmentUsecase(repo.Assignment),
		AuthN:      authNUsecase.NewAuthenticationUsecase(repo.AuthN, jwtService),
	}
}
