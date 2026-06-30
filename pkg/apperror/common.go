package apperror

import "net/http"

func New(
	status int,
	code string,
	message string,
	err ...error,
) *AppError {

	app := &AppError{
		Status:  status,
		Code:    code,
		Message: message,
	}

	if len(err) > 0 {
		app.Err = err[0]
	}

	return app
}

func BadRequest(message string, err ...error) *AppError {
	return New(http.StatusBadRequest, "BAD_REQUEST", message, err...)
}

func Unauthorized(message string, err ...error) *AppError {
	return New(http.StatusUnauthorized, "UNAUTHORIZED", message, err...)
}

func Forbidden(message string, err ...error) *AppError {
	return New(http.StatusForbidden, "FORBIDDEN", message, err...)
}

func NotFound(message string, err ...error) *AppError {
	return New(http.StatusNotFound, "NOT_FOUND", message, err...)
}

func Conflict(message string, err ...error) *AppError {
	return New(http.StatusConflict, "CONFLICT", message, err...)
}

func UnprocessableEntity(message string, err ...error) *AppError {
	return New(http.StatusUnprocessableEntity, "UNPROCESSABLE_ENTITY", message, err...)
}

func TooManyRequests(message string, err ...error) *AppError {
	return New(http.StatusTooManyRequests, "TOO_MANY_REQUESTS", message, err...)
}

func Internal(message string, err ...error) *AppError {
	return New(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", message, err...)
}
