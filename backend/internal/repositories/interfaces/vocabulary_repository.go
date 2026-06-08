package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type VocabularyRepository interface {
	Create(
		ctx context.Context,
		vocabulary *models.Vocabulary,
	) error

	FindByID(
		ctx context.Context,
		id uuid.UUID,
	) (*models.Vocabulary, error)

	FindByUserAndWord(
		ctx context.Context,
		userID uuid.UUID,
		word string,
	) (*models.Vocabulary, error)

	FindAllByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) ([]models.Vocabulary, error)

	Update(
		ctx context.Context,
		vocabulary *models.Vocabulary,
	) error

	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error
	SearchByUserID(
		ctx context.Context,
		userID uuid.UUID,
		search string,
	) ([]models.Vocabulary, error)
	FindRandomByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (*models.Vocabulary, error)
}