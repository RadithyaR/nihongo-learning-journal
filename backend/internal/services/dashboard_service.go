package services

import (
	"context"

	dashboardDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/dashboard"
	repositoryInterface "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterface "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/google/uuid"
)

type dashboardService struct {
	dashboardRepo repositoryInterface.DashboardRepository
}

func NewDashboardService(
	dashboardRepo repositoryInterface.DashboardRepository,
) serviceInterface.DashboardService {
	return &dashboardService{
		dashboardRepo: dashboardRepo,
	}
}

func (s *dashboardService) GetDashboard(
	ctx context.Context,
	userID uuid.UUID,
) (*dashboardDTO.DashboardResponse, error) {

	totalVocabulary, err := s.dashboardRepo.GetTotalVocabulary(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	totalKanji, err := s.dashboardRepo.GetTotalKanji(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	totalGrammar, err := s.dashboardRepo.GetTotalGrammar(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	reviewCountToday, err := s.dashboardRepo.GetReviewCountToday(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	activeGoals, completedGoals, err := s.dashboardRepo.GetGoalStatistics(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	studyStreak, err := s.dashboardRepo.GetStudyStreak(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	recentSessions, err := s.dashboardRepo.GetRecentSessions(
		ctx,
		userID,
		5,
	)
	if err != nil {
		return nil, err
	}

	sessionResponses := make(
		[]dashboardDTO.RecentSessionResponse,
		0,
		len(recentSessions),
	)

	for _, session := range recentSessions {

		sessionResponses = append(
			sessionResponses,
			dashboardDTO.RecentSessionResponse{
				ID:          session.ID,
				SessionDate: session.SessionDate,
				Notes:       session.Notes,
			},
		)
	}

	return &dashboardDTO.DashboardResponse{
		TotalVocabulary: int(totalVocabulary),
		TotalKanji:      int(totalKanji),
		TotalGrammar:    int(totalGrammar),

		ReviewCountToday: int(reviewCountToday),

		StudyStreak: studyStreak,

		ActiveGoals:    int(activeGoals),
		CompletedGoals: int(completedGoals),

		RecentSessions: sessionResponses,
	}, nil
}