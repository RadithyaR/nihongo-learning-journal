package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type EmailVerificationRepository interface {
	Create(
		ctx context.Context,
		verification *models.EmailVerification,
	) error

	FindByToken(
		ctx context.Context,
		token string,
	) (*models.EmailVerification, error)

	FindLatestByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (*models.EmailVerification, error)

	DeleteByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) error

	Update(
		ctx context.Context,
		verification *models.EmailVerification,
	) error
}