package services

import (
	"context"

	analyticsDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/analytics"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/google/uuid"
)

type analyticsService struct {
	analyticsRepository repositoryInterfaces.AnalyticsRepository
}

func NewAnalyticsService(
	analyticsRepository repositoryInterfaces.AnalyticsRepository,
) serviceInterfaces.AnalyticsService {
	return &analyticsService{
		analyticsRepository: analyticsRepository,
	}
}

func (s *analyticsService) GetAnalytics(
	ctx context.Context,
	userID uuid.UUID,
) (*analyticsDTO.AnalyticsResponse, error) {

	reviewsPerDay, err := s.analyticsRepository.GetReviewsPerDay(
		ctx,
		userID,
		30,
	)

	if err != nil {
		return nil, err
	}

	vocabularyGrowth, err := s.analyticsRepository.GetVocabularyGrowth(
		ctx,
		userID,
		30,
	)

	if err != nil {
		return nil, err
	}

	studySessions, err := s.analyticsRepository.GetStudySessionsPerDay(
		ctx,
		userID,
		30,
	)

	if err != nil {
		return nil, err
	}

	completionRate, err := s.analyticsRepository.GetGoalCompletionRate(
		ctx,
		userID,
	)

	if err != nil {
		return nil, err
	}

	return &analyticsDTO.AnalyticsResponse{
		ReviewsPerDay:    reviewsPerDay,
		VocabularyGrowth: vocabularyGrowth,
		StudySessions:    studySessions,
		CompletionRate:   completionRate,
	}, nil
}