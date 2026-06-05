package errors

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrSessionNotFound = errors.New("session not found")
	ErrInvalidGoogleToken = errors.New("invalid google token")
	ErrEmailNotVerified = errors.New("email not verified")
)