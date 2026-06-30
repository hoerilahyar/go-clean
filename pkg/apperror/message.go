package apperror

var (
	// Authentication
	ErrInvalidCredential = Unauthorized("Invalid username/email or password")
	ErrUserInactive      = Forbidden("User account is inactive")

	// Token
	ErrGenerateAccessToken  = Internal("Failed to generate access token")
	ErrGenerateRefreshToken = Internal("Failed to generate refresh token")
	ErrInvalidRefreshToken  = Unauthorized("Invalid refresh token")
	ErrRefreshTokenExpired  = Unauthorized("Refresh token has expired")
	ErrDeleteResetToken     = Internal("Failed to delete password reset token")

	// Session
	ErrCreateSession   = Internal("Failed to create session")
	ErrUpdateSession   = Internal("Failed to update session")
	ErrDeleteSession   = Internal("Failed to delete session")
	ErrSessionNotFound = Unauthorized("Session not found")

	// Password
	ErrCurrentPasswordIncorrect     = BadRequest("Current password is incorrect")
	ErrPasswordConfirmationMismatch = BadRequest("Password confirmation does not match")
	ErrPasswordMustBeDifferent      = BadRequest("New password must be different from current password")
	ErrHashPassword                 = Internal("Failed to hash password")
	ErrUpdatePassword               = Internal("Failed to update password")

	// Reset Password
	ErrGenerateResetToken = Internal("Failed to generate password reset token")
	ErrCreateResetToken   = Internal("Failed to create password reset token")
	ErrResetTokenNotFound = BadRequest("Invalid reset password token")
	ErrResetTokenExpired  = BadRequest("Reset password token has expired")

	// User
	ErrUserNotFound = NotFound("User not found")
)
