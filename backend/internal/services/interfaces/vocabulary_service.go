package interfaces

import (
	"context"

	vocabularyDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/vocabulary"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type VocabularyService interface {
	CreateVocabulary(
		ctx context.Context,
		userID uuid.UUID,
		dto vocabularyDTO.CreateVocabularyRequest,
	) (*vocabularyDTO.VocabularyResponse, error)

	GetVocabularyByID(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
	) (*vocabularyDTO.VocabularyResponse, error)

	GetVocabularyList(
		ctx context.Context,
		userID uuid.UUID,
		filter models.ListFilter,
	) ([]vocabularyDTO.VocabularyResponse, error)

	UpdateVocabulary(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
		dto vocabularyDTO.UpdateVocabularyRequest,
	) (*vocabularyDTO.VocabularyResponse, error)

	DeleteVocabulary(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
	) error
	ToggleFavourite(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
	) (*vocabularyDTO.VocabularyResponse, error)
}