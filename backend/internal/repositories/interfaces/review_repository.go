package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type ReviewRepository interface {
	Create(
		ctx context.Context,
		review *models.ReviewLog,
	) error

	GetReviewCountByUser(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)
}