package interfaces

import (
	"context"

	kanjiDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/kanji"
	"github.com/google/uuid"
)

type KanjiService interface {
	CreateKanji(
		ctx context.Context,
		userID uuid.UUID,
		dto kanjiDTO.CreateKanjiRequest,
	) (*kanjiDTO.KanjiResponse, error)

	GetKanjiByID(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
	) (*kanjiDTO.KanjiResponse, error)

	GetKanjiList(
		ctx context.Context,
		userID uuid.UUID,
		search string,
	) ([]kanjiDTO.KanjiResponse, error)

	UpdateKanji(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
		dto kanjiDTO.UpdateKanjiRequest,
	) (*kanjiDTO.KanjiResponse, error)

	DeleteKanji(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
	) error

	ToggleFavourite(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
	) (*kanjiDTO.KanjiResponse, error)
}