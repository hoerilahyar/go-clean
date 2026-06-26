package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/response"
	authenticationRepo "github.com/hoerilahyar/go-clean/internal/domain/authentication/repository"
	"github.com/hoerilahyar/go-clean/pkg/jwt"
)

type AuthenticationUsecase interface {
	Login(ctx context.Context, req request.LoginRequest) (*response.LoginResponse, error)

	RefreshToken(ctx context.Context, req request.RefreshTokenRequest) (*response.RefreshTokenResponse, error)

	Logout(ctx context.Context, req request.LogoutRequest) error

	ForgotPassword(ctx context.Context, req request.ForgotPasswordRequest) error

	ResetPassword(ctx context.Context, req request.ResetPasswordRequest) error

	ChangePassword(
		ctx context.Context,
		userID uint64,
		req request.ChangePasswordRequest,
	) error
}

type authenticationUsecase struct {
	repository authenticationRepo.AuthenticationRepository
	jwtService jwt.Service
}

func NewAuthenticationUsecase(
	repository authenticationRepo.AuthenticationRepository,
	jwtService jwt.Service,
) AuthenticationUsecase {
	return &authenticationUsecase{
		repository: repository,
		jwtService: jwtService,
	}
}
