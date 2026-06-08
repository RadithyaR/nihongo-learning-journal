package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type GrammarRepository interface {
	Create(
		ctx context.Context,
		grammar *models.Grammar,
	) error

	FindByID(
		ctx context.Context,
		id uuid.UUID,
	) (*models.Grammar, error)

	FindByUserAndPattern(
		ctx context.Context,
		userID uuid.UUID,
		pattern string,
	) (*models.Grammar, error)

	FindAllByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) ([]models.Grammar, error)

	SearchByUserID(
		ctx context.Context,
		userID uuid.UUID,
		search string,
	) ([]models.Grammar, error)

	Update(
		ctx context.Context,
		grammar *models.Grammar,
	) error

	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error

	FindRandomByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (*models.Grammar, error)
}