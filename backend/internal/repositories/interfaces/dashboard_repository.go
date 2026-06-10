package interfaces

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	"github.com/google/uuid"
)

type DashboardRepository interface {
	GetTotalVocabulary(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)

	GetTotalKanji(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)

	GetTotalGrammar(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)

	GetReviewCountToday(
		ctx context.Context,
		userID uuid.UUID,
	) (int64, error)

	GetGoalStatistics(
		ctx context.Context,
		userID uuid.UUID,
	) (
		active int64,
		completed int64,
		err error,
	)

	GetRecentSessions(
		ctx context.Context,
		userID uuid.UUID,
		limit int,
	) ([]models.StudySession, error)

	GetStudyStreak(
		ctx context.Context,
		userID uuid.UUID,
	) (int, error)
}