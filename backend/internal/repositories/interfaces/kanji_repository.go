package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type KanjiRepository interface {
	Create(
		ctx context.Context,
		kanji *models.Kanji,
	) error

	FindByID(
		ctx context.Context,
		id uuid.UUID,
	) (*models.Kanji, error)

	FindByUserAndCharacter(
		ctx context.Context,
		userID uuid.UUID,
		character string,
	) (*models.Kanji, error)

	FindAllByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) ([]models.Kanji, error)

	SearchByUserID(
		ctx context.Context,
		userID uuid.UUID,
		search string,
	) ([]models.Kanji, error)

	Update(
		ctx context.Context,
		kanji *models.Kanji,
	) error

	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error
	FindRandomByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (*models.Kanji, error)
}