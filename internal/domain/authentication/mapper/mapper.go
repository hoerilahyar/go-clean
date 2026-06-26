package mapper

import (
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/response"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/entity"
	userEntity "github.com/hoerilahyar/go-clean/internal/domain/user/entity"
)

func ToLoginResponse(
	user userEntity.User,
	token entity.Token,
) *response.LoginResponse {

	return &response.LoginResponse{
		User: response.UserResponse{
			ID:       user.ID,
			FullName: user.FullName,
			Username: user.Username,
			Email:    user.Email,
			Status:   user.Status,
		},
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		ExpiresIn:    token.ExpiresIn,
	}
}

func ToRefreshTokenResponse(
	token entity.Token,
) *response.RefreshTokenResponse {

	return &response.RefreshTokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		ExpiresIn:    token.ExpiresIn,
	}
}
