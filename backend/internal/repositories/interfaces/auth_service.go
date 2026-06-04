package interfaces

import (
	"context"

	authDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/auth"
	responses "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
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
}

