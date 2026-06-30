package usecase

import (
	"context"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/pkg/apperror"
	"golang.org/x/crypto/bcrypt"
)

func (u *authenticationUsecase) ChangePassword(
	ctx context.Context,
	userID uint64,
	req request.ChangePasswordRequest,
) error {

	// Retrieve user by ID.
	user, err := u.repository.FindUserByID(
		ctx,
		userID,
	)
	if err != nil {
		return err
	}

	// Ensure the user exists.
	if user == nil {
		return apperror.ErrUserNotFound
	}

	// Verify the current password.
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.OldPassword),
	); err != nil {
		return apperror.ErrCurrentPasswordIncorrect
	}

	// Ensure the new password is different from the current password.
	if req.OldPassword == req.NewPassword {
		return apperror.ErrPasswordMustBeDifferent
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
		userID,
		string(hashedPassword),
	); err != nil {
		return err
	}

	// Revoke all active sessions after password change.
	if err := u.repository.DeleteSessionsByUserID(
		ctx,
		userID,
	); err != nil {
		return err
	}

	return nil
}
