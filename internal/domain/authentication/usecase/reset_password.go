package usecase

import (
	"context"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	authErr "github.com/hoerilahyar/go-clean/internal/domain/authentication/errors"
	"golang.org/x/crypto/bcrypt"
)

func (u *authenticationUsecase) ResetPassword(
	ctx context.Context,
	req request.ResetPasswordRequest,
) error {

	reset, err := u.repository.FindPasswordResetToken(
		ctx,
		req.Token,
	)
	if err != nil {
		return err
	}

	if time.Now().After(reset.ExpiredAt) {
		// return ErrResetTokenExpired
		return authErr.ErrRefreshTokenExpired
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	if err := u.repository.UpdatePassword(
		ctx,
		reset.UserID,
		string(hashedPassword),
	); err != nil {
		return err
	}

	if err := u.repository.DeleteSessionsByUserID(
		ctx,
		reset.UserID,
	); err != nil {
		return err
	}

	if err := u.repository.DeletePasswordResetToken(
		ctx,
		req.Token,
	); err != nil {
		return err
	}

	return nil
}
