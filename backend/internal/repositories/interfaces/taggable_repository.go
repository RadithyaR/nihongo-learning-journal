package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type TaggableRepository interface {
	Create(
		ctx context.Context,
		taggable *models.Taggable,
	) error

	Delete(
		ctx context.Context,
		tagID uuid.UUID,
		itemType string,
		itemID uuid.UUID,
	) error

	Exists(
		ctx context.Context,
		tagID uuid.UUID,
		itemType string,
		itemID uuid.UUID,
	) (bool, error)

	FindTagsByItem(
		ctx context.Context,
		itemType string,
		itemID uuid.UUID,
	) ([]models.Tag, error)
}