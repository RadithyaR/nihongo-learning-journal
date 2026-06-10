package interfaces

import (
	"context"

	grammarDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/grammar"
	"github.com/google/uuid"
)

type GrammarService interface {
	CreateGrammar(
		ctx context.Context,
		userID uuid.UUID,
		dto grammarDTO.CreateGrammarRequest,
	) (*grammarDTO.GrammarResponse, error)

	GetGrammarByID(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
	) (*grammarDTO.GrammarResponse, error)

	GetGrammarList(
		ctx context.Context,
		userID uuid.UUID,
		search string,
	) ([]grammarDTO.GrammarResponse, error)

	UpdateGrammar(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
		dto grammarDTO.UpdateGrammarRequest,
	) (*grammarDTO.GrammarResponse, error)

	DeleteGrammar(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
	) error

	ToggleFavourite(
		ctx context.Context,
		userID uuid.UUID,
		id uuid.UUID,
	) (*grammarDTO.GrammarResponse, error)
}