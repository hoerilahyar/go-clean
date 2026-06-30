package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func (u *authenticationUsecase) Logout(
	ctx context.Context,
	req request.LogoutRequest,
) error {

	// Retrieve the current session.
	session, err := u.repository.FindSessionByRefreshToken(
		ctx,
		req.RefreshToken,
	)
	if err != nil {
		return err
	}

	// Ensure the session exists.
	if session == nil {
		return apperror.ErrSessionNotFound
	}

	// Delete the current session.
	if err := u.repository.DeleteSessionByRefreshToken(
		ctx,
		req.RefreshToken,
	); err != nil {
		return err
	}

	return nil
}
