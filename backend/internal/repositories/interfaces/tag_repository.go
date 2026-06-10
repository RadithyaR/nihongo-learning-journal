package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type TagRepository interface {
	Create(
		ctx context.Context,
		tag *models.Tag,
	) error

	FindByID(
		ctx context.Context,
		id uuid.UUID,
	) (*models.Tag, error)

	FindByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) ([]models.Tag, error)

	Update(
		ctx context.Context,
		tag *models.Tag,
	) error

	Delete(
		ctx context.Context,
		tag *models.Tag,
	) error

	ExistsByName(
		ctx context.Context,
		userID uuid.UUID,
		name string,
	) (bool, error)
}