package usecase

import (
	"context"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/response"

	authErr "github.com/hoerilahyar/go-clean/internal/domain/authentication/errors"
)

func (u *authenticationUsecase) RefreshToken(
	ctx context.Context,
	req request.RefreshTokenRequest,
) (*response.RefreshTokenResponse, error) {

	// Cari session
	session, err := u.repository.FindSessionByRefreshToken(
		ctx,
		req.RefreshToken,
	)
	if err != nil {
		return nil, err
	}

	// Cek expired
	if time.Now().After(session.ExpiredAt) {
		return nil, authErr.ErrRefreshTokenExpired
	}

	// Ambil user
	user, err := u.repository.FindUserByID(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	// Generate Access Token baru
	accessToken, err := u.jwtService.GenerateAccessToken(
		user.ID,
		user.Username,
	)
	if err != nil {
		return nil, err
	}

	// Generate Refresh Token baru
	refreshToken, expiredAt, err := u.jwtService.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Update session
	session.RefreshToken = refreshToken
	session.ExpiredAt = expiredAt

	if err := u.repository.UpdateSession(ctx, *session); err != nil {
		return nil, err
	}

	return &response.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    u.jwtService.AccessTokenTTL(),
	}, nil
}
