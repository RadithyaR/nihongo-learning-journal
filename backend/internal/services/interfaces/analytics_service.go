package interfaces

import (
	"context"

	analyticsDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/analytics"
	"github.com/google/uuid"
)

type AnalyticsService interface {
	GetAnalytics(
		ctx context.Context,
		userID uuid.UUID,
	) (*analyticsDTO.AnalyticsResponse, error)
}

