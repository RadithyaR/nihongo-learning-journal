package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type SRSService interface {

	UpdateReview(
		ctx context.Context,
		userID uuid.UUID,
		itemType string,
		itemID uuid.UUID,
		rating string,
	) error

	GetDueCount(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)
	GetOverdueCount(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)
	GetNextDueItem(
		ctx context.Context,
		userID uuid.UUID,
		itemType string,
	) (*models.SRSRecord, error)
}