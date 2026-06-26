package handler

import "github.com/hoerilahyar/go-clean/internal/domain/authentication/usecase"

type AuthenticationHandler struct {
	usecase usecase.AuthenticationUsecase
}

func NewAuthenticationHandler(
	usecase usecase.AuthenticationUsecase,
) *AuthenticationHandler {
	return &AuthenticationHandler{
		usecase: usecase,
	}
}
