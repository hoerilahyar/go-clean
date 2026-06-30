package usecase

import (
	"context"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/response"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func (u *authenticationUsecase) RefreshToken(
	ctx context.Context,
	req request.RefreshTokenRequest,
) (*response.RefreshTokenResponse, error) {

	// Retrieve the session by refresh token.
	session, err := u.repository.FindSessionByRefreshToken(
		ctx,
		req.RefreshToken,
	)
	if err != nil {
		return nil, err
	}

	// Ensure the session exists.
	if session == nil {
		return nil, apperror.ErrSessionNotFound
	}

	// Validate refresh token expiration.
	if time.Now().After(session.ExpiredAt) {
		return nil, apperror.ErrRefreshTokenExpired
	}

	// Retrieve the user.
	user, err := u.repository.FindUserByID(
		ctx,
		session.UserID,
	)
	if err != nil {
		return nil, err
	}

	// Ensure the user exists.
	if user == nil {
		return nil, apperror.ErrUserNotFound
	}

	// Ensure the user account is active.
	if user.Status != "ACTIVE" {
		return nil, apperror.ErrUserInactive
	}

	// Generate a new access token.
	accessToken, err := u.jwtService.GenerateAccessToken(
		user.ID,
		user.Username,
	)
	if err != nil {
		return nil, apperror.Internal("Failed to generate access token", err)
	}

	// Generate a new refresh token.
	refreshToken, expiredAt, err := u.jwtService.GenerateRefreshToken()
	if err != nil {
		return nil, apperror.Internal("Failed to generate refresh token", err)
	}

	// Update the current session.
	session.RefreshToken = refreshToken
	session.ExpiredAt = expiredAt

	if err := u.repository.UpdateSession(
		ctx,
		*session,
	); err != nil {
		return nil, err
	}

	// Build the refresh token response.
	return &response.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    u.jwtService.AccessTokenTTL(),
	}, nil
}
