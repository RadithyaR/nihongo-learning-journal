package interfaces

import (
	"context"

	reviewDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/review"
	"github.com/google/uuid"
)

type ReviewService interface {
	GetNextReview(
		ctx context.Context,
		userID uuid.UUID,
	) (*reviewDTO.NextReviewResult, error)

	SubmitReview(
		ctx context.Context,
		userID uuid.UUID,
		dto reviewDTO.SubmitReviewRequest,
	) error
	GetNextKanjiReview(
		ctx context.Context,
		userID uuid.UUID,
	) (*reviewDTO.NextKanjiReviewResult, error)

	SubmitKanjiReview(
		ctx context.Context,
		userID uuid.UUID,
		dto reviewDTO.SubmitReviewRequest,
	) error
	GetNextGrammarReview(
		ctx context.Context,
		userID uuid.UUID,
	) (*reviewDTO.NextGrammarReviewResult, error)

	SubmitGrammarReview(
		ctx context.Context,
		userID uuid.UUID,
		dto reviewDTO.SubmitGrammarReviewRequest,
	) error
}