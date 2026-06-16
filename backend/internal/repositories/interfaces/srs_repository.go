package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type SRSRepository interface {

	Create(
		ctx context.Context,
		record *models.SRSRecord,
	) error

	Update(
		ctx context.Context,
		record *models.SRSRecord,
	) error

	FindByItem(
		ctx context.Context,
		userID uuid.UUID,
		itemType string,
		itemID uuid.UUID,
	) (*models.SRSRecord, error)

	GetDueItems(
		ctx context.Context,
		userID uuid.UUID,
	) ([]models.SRSRecord, error)

	CountDueToday(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)
	CountOverdue(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)
	GetNextDueItem(
		ctx context.Context,
		userID uuid.UUID,
		itemType string,
	) (*models.SRSRecord, error)
}