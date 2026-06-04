package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type UserSessionRepository interface {
	Create(
		ctx context.Context,
		session *models.UserSession,
	) error
	FindByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (*models.UserSession, error)
	FindByTokenID(
		ctx context.Context,
		tokenID uuid.UUID,
	) (*models.UserSession, error)
	Update(
		ctx context.Context,
		session *models.UserSession,
	) error

	Delete(
		ctx context.Context,
		sessionID uuid.UUID,
	) error
	DeleteByTokenID(
		ctx context.Context,
		tokenID uuid.UUID,
	) error
}