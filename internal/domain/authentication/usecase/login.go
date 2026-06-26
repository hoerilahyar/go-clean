package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/entity"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/mapper"

	"golang.org/x/crypto/bcrypt"
)

func (u *authenticationUsecase) Login(
	ctx context.Context,
	req request.LoginRequest,
) (*response.LoginResponse, error) {

	// Cari user berdasarkan username/email
	user, err := u.repository.FindUserByIdentity(ctx, req.Identity)
	if err != nil {
		return nil, err
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return nil, err
	}

	// Generate Access Token
	accessToken, err := u.jwtService.GenerateAccessToken(
		user.ID,
		user.Username,
	)
	if err != nil {
		return nil, err
	}

	// Generate Refresh Token
	refreshToken, expiredAt, err := u.jwtService.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Simpan session
	session := entity.Session{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}

	if err := u.repository.CreateSession(ctx, session); err != nil {
		return nil, err
	}

	token := entity.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    u.jwtService.AccessTokenTTL(),
	}

	return mapper.ToLoginResponse(*user, token), nil
}
