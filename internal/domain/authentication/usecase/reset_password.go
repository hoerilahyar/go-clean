package usecase

import (
	"context"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
	"golang.org/x/crypto/bcrypt"
)

func (u *authenticationUsecase) ResetPassword(
	ctx context.Context,
	req request.ResetPasswordRequest,
) error {

	// Retrieve the password reset token.
	reset, err := u.repository.FindPasswordResetToken(
		ctx,
		req.Token,
	)
	if err != nil {
		return err
	}

	// Ensure the reset token exists.
	if reset == nil {
		return apperror.ErrResetTokenNotFound
	}

	// Validate reset token expiration.
	if time.Now().After(reset.ExpiredAt) {

		// Remove the expired reset token.
		_ = u.repository.DeletePasswordResetToken(
			ctx,
			req.Token,
		)

		return apperror.ErrResetTokenExpired
	}

	// Validate password confirmation.
	if req.NewPassword != req.ConfirmPassword {
		return apperror.ErrPasswordConfirmationMismatch
	}

	// Generate a new password hash.
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return apperror.Internal("Failed to hash password", err)
	}

	// Update the user's password.
	if err := u.repository.UpdatePassword(
		ctx,
		reset.UserID,
		string(hashedPassword),
	); err != nil {
		return err
	}

	// Revoke all active sessions.
	if err := u.repository.DeleteSessionsByUserID(
		ctx,
		reset.UserID,
	); err != nil {
		return err
	}

	// Remove the used reset token.
	if err := u.repository.DeletePasswordResetToken(
		ctx,
		req.Token,
	); err != nil {
		return err
	}

	return nil
}
