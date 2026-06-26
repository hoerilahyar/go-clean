package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	authErr "github.com/hoerilahyar/go-clean/internal/domain/authentication/errors"
	"golang.org/x/crypto/bcrypt"
)

func (u *authenticationUsecase) ChangePassword(
	ctx context.Context,
	userID uint64,
	req request.ChangePasswordRequest,
) error {

	// Ambil user
	user, err := u.repository.FindUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// Verifikasi password lama
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.OldPassword),
	); err != nil {
		return err
	}

	// Validasi konfirmasi password
	if req.NewPassword != req.ConfirmPassword {
		return authErr.ErrPasswordConfirmationMismatch
	}

	// Hash password baru
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	// Update password
	if err := u.repository.UpdatePassword(
		ctx,
		userID,
		string(hashedPassword),
	); err != nil {
		return err
	}

	// Logout semua device
	if err := u.repository.DeleteSessionsByUserID(
		ctx,
		userID,
	); err != nil {
		return err
	}

	return nil
}
