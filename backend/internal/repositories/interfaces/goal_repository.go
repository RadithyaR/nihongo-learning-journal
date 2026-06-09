package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type GoalRepository interface {
	Create(
		ctx context.Context,
		goal *models.Goal,
	) error

	FindByID(
		ctx context.Context,
		id uuid.UUID,
	) (*models.Goal, error)

	FindAllByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) ([]models.Goal, error)

	Update(
		ctx context.Context,
		goal *models.Goal,
	) error

	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error
	CountProgress(
		ctx context.Context,
		userID uuid.UUID,
		goalType string,
	) (int64, error)
}