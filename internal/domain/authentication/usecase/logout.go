package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
)

func (u *authenticationUsecase) Logout(
	ctx context.Context,
	req request.LogoutRequest,
) error {

	// Pastikan session masih ada
	_, err := u.repository.FindSessionByRefreshToken(
		ctx,
		req.RefreshToken,
	)
	if err != nil {
		return err
	}

	// Hapus session
	if err := u.repository.DeleteSessionByRefreshToken(
		ctx,
		req.RefreshToken,
	); err != nil {
		return err
	}

	return nil
}
