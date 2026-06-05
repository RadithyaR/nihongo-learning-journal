package interfaces

import (
	"context"

	authDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/auth"
	responses "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(
		ctx context.Context,
		dto authDTO.RegisterDTO,
	) (*responses.UserResponse, error)

	Login(
		ctx context.Context,
		dto authDTO.LoginDTO,
	) (*responses.LoginResponse, error)
	RefreshToken(
		ctx context.Context,
		refreshToken string,
	) (*responses.LoginResponse, error)
	Logout(
		ctx context.Context,
		refreshToken string,
	) error
	LogoutAll(
		ctx context.Context,
		userID uuid.UUID,
	) error
	VerifyEmail(
		ctx context.Context,
		token string,
	) error
	ForgotPassword(
		ctx context.Context,
		dto authDTO.ForgotPasswordDTO,
	) (string, error)
	ResetPassword(
		ctx context.Context,
		dto authDTO.ResetPasswordDTO,
	) error
	ChangePassword(
		ctx context.Context,
		userID uuid.UUID,
		dto authDTO.ChangePasswordDTO,
	) error
}

