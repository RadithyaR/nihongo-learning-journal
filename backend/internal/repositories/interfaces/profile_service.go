package interfaces

import (
	"context"

	profileDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/profile"
	"github.com/google/uuid"
)

type ProfileService interface {
	GetProfile(
		ctx context.Context,
		userID uuid.UUID,
	) (*profileDTO.ProfileResponse, error)

	UpdateProfile(
		ctx context.Context,
		userID uuid.UUID,
		dto profileDTO.UpdateProfileRequest,
	) (*profileDTO.ProfileResponse, error)
}