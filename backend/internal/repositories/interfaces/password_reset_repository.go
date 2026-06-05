package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type PasswordResetRepository interface {
	Create(
		ctx context.Context,
		reset *models.PasswordReset,
	) error

	FindByToken(
		ctx context.Context,
		token string,
	) (*models.PasswordReset, error)

	FindLatestByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (*models.PasswordReset, error)

	Update(
		ctx context.Context,
		reset *models.PasswordReset,
	) error

	DeleteByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) error
}