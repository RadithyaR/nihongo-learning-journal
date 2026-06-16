package interfaces

import (
	"context"

	analyticsDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/analytics"
	"github.com/google/uuid"
)

type AnalyticsRepository interface {
	GetReviewsPerDay(
		ctx context.Context,
		userID uuid.UUID,
		days int,
	) ([]analyticsDTO.DailyStat, error)

	GetVocabularyGrowth(
		ctx context.Context,
		userID uuid.UUID,
		days int,
	) ([]analyticsDTO.DailyStat, error)

	GetStudySessionsPerDay(
		ctx context.Context,
		userID uuid.UUID,
		days int,
	) ([]analyticsDTO.DailyStat, error)

	GetGoalCompletionRate(
		ctx context.Context,
		userID uuid.UUID,
	) (float64, error)
}