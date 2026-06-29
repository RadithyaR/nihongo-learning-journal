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

	FindFiltered(
		ctx context.Context,
		userID uuid.UUID,
		filter models.ListFilter,
	) ([]models.Vocabulary, error)

	Update(
		ctx context.Context,
		vocabulary *models.Vocabulary,
	) error

	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error

	DeleteMeaningsByVocabularyID(
		ctx context.Context,
		vocabularyID uuid.UUID,
	) error

	CreateMeanings(
		ctx context.Context,
		meanings []models.VocabularyMeaning,
	) error

	FindRandomByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (*models.Vocabulary, error)
}