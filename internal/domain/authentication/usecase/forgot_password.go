package usecase

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/hoerilahyar/go-clean/internal/domain/authentication/dto/request"
	"github.com/hoerilahyar/go-clean/internal/domain/authentication/entity"
)

func (u *authenticationUsecase) ForgotPassword(
	ctx context.Context,
	req request.ForgotPasswordRequest,
) error {

	// Cari user berdasarkan email
	user, err := u.repository.FindUserByIdentity(ctx, req.Email)
	if err != nil {
		return err
	}

	// Hapus token lama (abaikan jika tidak ada)
	_ = u.repository.DeletePasswordResetTokenByUserID(
		ctx,
		user.ID,
	)

	// Generate token random
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return err
	}

	token := hex.EncodeToString(b)

	reset := entity.PasswordReset{
		UserID:    user.ID,
		Token:     token,
		ExpiredAt: time.Now().Add(15 * time.Minute),
	}

	if err := u.repository.CreatePasswordResetToken(
		ctx,
		reset,
	); err != nil {
		return err
	}

	// TODO:
	// Kirim email berisi link reset password
	//
	// https://example.com/reset-password?token=xxxxx

	return nil
}
