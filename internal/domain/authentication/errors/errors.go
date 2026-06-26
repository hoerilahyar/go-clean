package errors

import "errors"

var (
	ErrInvalidCredential            = errors.New("invalid username or password")
	ErrRefreshTokenExpired          = errors.New("refresh token expired")
	ErrInvalidRefreshToken          = errors.New("invalid refresh token")
	ErrAccountInactive              = errors.New("account is inactive")
	ErrResetTokenExpired            = errors.New("reset token expired")
	ErrPasswordConfirmationMismatch = errors.New("password confirmation mismatch")
)
