package usecase

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/entity"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
)

func (u *authenticationUsecase) ForgotPassword(
	ctx context.Context,
	req request.ForgotPasswordRequest,
) error {

	// Retrieve the user by email.
	user, err := u.repository.FindUserByIdentity(
		ctx,
		req.Email,
	)
	if err != nil {
		return err
	}

	// Always return success to prevent user enumeration.
	if user == nil {
		return nil
	}

	// Remove any existing reset token.
	// Ignore errors because this is only a cleanup step.
	_ = u.repository.DeletePasswordResetTokenByUserID(
		ctx,
		user.ID,
	)

	// Generate a secure random reset token.
	buffer := make([]byte, 32)

	if _, err := rand.Read(buffer); err != nil {
		return apperror.Internal("Failed to generate password reset token", err)
	}

	token := hex.EncodeToString(buffer)

	reset := entity.PasswordReset{
		UserID:    user.ID,
		Token:     token,
		ExpiredAt: time.Now().Add(15 * time.Minute),
	}

	// Store the password reset token.
	if err := u.repository.CreatePasswordResetToken(
		ctx,
		reset,
	); err != nil {
		return err
	}

	// Send the password reset email.
	//
	// TODO:
	// u.mailService.SendResetPasswordEmail(user.Email, token)

	return nil
}
