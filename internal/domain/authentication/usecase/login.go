package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/mapper"
	"github.com/hoerilahyar/go-clean/pkg/apperror"

	"golang.org/x/crypto/bcrypt"
)

func (u *authenticationUsecase) Login(
	ctx context.Context,
	req request.LoginRequest,
) (*response.LoginResponse, error) {

	// Retrieve user by username or email.
	user, err := u.repository.FindUserByIdentity(
		ctx,
		req.Identity,
	)
	if err != nil {
		return nil, err
	}

	// Prevent user enumeration.
	if user == nil {
		return nil, apperror.ErrInvalidCredential
	}

	// Verify the password.
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return nil, apperror.ErrInvalidCredential
	}

	// Ensure the user account is active.
	if user.Status != "ACTIVE" {
		return nil, apperror.ErrUserInactive
	}

	// Generate access token.
	accessToken, err := u.jwtService.GenerateAccessToken(
		user.ID,
		user.Username,
	)
	if err != nil {
		return nil, apperror.Internal("Failed to generate access token", err)
	}

	// Generate refresh token.
	refreshToken, expiredAt, err := u.jwtService.GenerateRefreshToken()
	if err != nil {
		return nil, apperror.Internal("Failed to generate refresh token", err)
	}

	// Create user session.
	session := entity.Session{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}

	if err := u.repository.CreateSession(
		ctx,
		session,
	); err != nil {
		return nil, err
	}

	// Build token payload.
	token := entity.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    u.jwtService.AccessTokenTTL(),
	}

	return mapper.ToLoginResponse(*user, token), nil
}
