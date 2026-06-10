package repositories

import (
	"context"
	"time"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/constants"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterface "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(
	db *gorm.DB,
) repositoryInterface.DashboardRepository {
	return &dashboardRepository{
		db: db,
	}
}

func (r *dashboardRepository) GetTotalVocabulary(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {

	var count int64

	err := r.db.WithContext(ctx).
		Model(&models.Vocabulary{}).
		Where("user_id = ?", userID).
		Count(&count).Error

	return count, err
}

func (r *dashboardRepository) GetTotalKanji(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {

	var count int64

	err := r.db.WithContext(ctx).
		Model(&models.Kanji{}).
		Where("user_id = ?", userID).
		Count(&count).Error

	return count, err
}

func (r *dashboardRepository) GetTotalGrammar(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {

	var count int64

	err := r.db.WithContext(ctx).
		Model(&models.Grammar{}).
		Where("user_id = ?", userID).
		Count(&count).Error

	return count, err
}

func (r *dashboardRepository) GetReviewCountToday(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {

	var count int64

	now := time.Now()

	startOfDay := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	endOfDay := startOfDay.Add(24 * time.Hour)

	err := r.db.WithContext(ctx).
		Model(&models.ReviewLog{}).
		Where("user_id = ?", userID).
		Where("reviewed_at >= ?", startOfDay).
		Where("reviewed_at < ?", endOfDay).
		Count(&count).Error

	return count, err
}

func (r *dashboardRepository) GetGoalStatistics(
	ctx context.Context,
	userID uuid.UUID,
) (
	active int64,
	completed int64,
	err error,
) {

	err = r.db.WithContext(ctx).
		Model(&models.Goal{}).
		Where(
			"user_id = ? AND status = ?",
			userID,
			constants.GoalStatusInProgress,
		).
		Count(&active).Error

	if err != nil {
		return
	}

	err = r.db.WithContext(ctx).
		Model(&models.Goal{}).
		Where(
			"user_id = ? AND status = ?",
			userID,
			constants.GoalStatusCompleted,
		).
		Count(&completed).Error

	return
}

func (r *dashboardRepository) GetRecentSessions(
	ctx context.Context,
	userID uuid.UUID,
	limit int,
) ([]models.StudySession, error) {

	var sessions []models.StudySession

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("session_date DESC").
		Limit(limit).
		Find(&sessions).Error

	return sessions, err
}

func (r *dashboardRepository) GetStudyStreak(
	ctx context.Context,
	userID uuid.UUID,
) (int, error) {

	var sessions []models.StudySession

	err := r.db.WithContext(ctx).
		Select("session_date").
		Where("user_id = ?", userID).
		Order("session_date DESC").
		Find(&sessions).Error

	if err != nil {
		return 0, err
	}

	if len(sessions) == 0 {
		return 0, nil
	}

	now := time.Now()

	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		now.Location(),
	)

	latestSession := sessions[0].SessionDate

	latestSessionDate := time.Date(
		latestSession.Year(),
		latestSession.Month(),
		latestSession.Day(),
		0,
		0,
		0,
		0,
		latestSession.Location(),
	)

	// Jika belum belajar hari ini maka streak = 0
	if !latestSessionDate.Equal(today) {
		return 0, nil
	}

	streak := 1

	for i := 1; i < len(sessions); i++ {

		previousDate := time.Date(
			sessions[i-1].SessionDate.Year(),
			sessions[i-1].SessionDate.Month(),
			sessions[i-1].SessionDate.Day(),
			0,
			0,
			0,
			0,
			sessions[i-1].SessionDate.Location(),
		)

		currentDate := time.Date(
			sessions[i].SessionDate.Year(),
			sessions[i].SessionDate.Month(),
			sessions[i].SessionDate.Day(),
			0,
			0,
			0,
			0,
			sessions[i].SessionDate.Location(),
		)

		expectedPreviousDay := previousDate.AddDate(0, 0, -1)

		if currentDate.Equal(expectedPreviousDay) {
			streak++
			continue
		}

		break
	}

	return streak, nil
}