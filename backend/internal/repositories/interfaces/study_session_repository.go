package interfaces

import (
	"context"
	"time"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type StudySessionRepository interface {
	Create(
		ctx context.Context,
		session *models.StudySession,
	) error

	FindByUserAndDate(
		ctx context.Context,
		userID uuid.UUID,
		sessionDate time.Time,
	) (*models.StudySession, error)

	Update(
		ctx context.Context,
		session *models.StudySession,
	) error

	AddItem(
		ctx context.Context,
		item *models.StudySessionItem,
	) error

	ItemExists(
		ctx context.Context,
		studySessionID uuid.UUID,
		itemType string,
		itemID uuid.UUID,
	) (bool, error)

	FindHistoryByUserID(
		ctx context.Context,
		userID uuid.UUID,
		page int,
		limit int,
	) ([]models.StudySession, int64, error)
	FindItemsBySessionID(
		ctx context.Context,
		sessionID uuid.UUID,
	) ([]models.StudySessionItem, error)
}