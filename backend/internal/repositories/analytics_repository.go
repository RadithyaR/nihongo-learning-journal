package repositories

import (
	"context"
	"time"

	analyticsDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/analytics"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type analyticsRepository struct {
	db *gorm.DB
}

func NewAnalyticsRepository(
	db *gorm.DB,
) repositoryInterfaces.AnalyticsRepository {
	return &analyticsRepository{
		db: db,
	}
}

type dailyStatResult struct {
	Date  string
	Count int
}

func (r *analyticsRepository) GetReviewsPerDay(
	ctx context.Context,
	userID uuid.UUID,
	days int,
) ([]analyticsDTO.DailyStat, error) {

	var results []dailyStatResult

	startDate := time.Now().AddDate(
		0,
		0,
		-days,
	)

	err := r.db.WithContext(
		ctx,
	).Table(
		"review_logs",
	).Select(
		"DATE(reviewed_at) as date, COUNT(*) as count",
	).Where(
		"user_id = ?",
		userID,
	).Where(
		"reviewed_at >= ?",
		startDate,
	).Group(
		"DATE(reviewed_at)",
	).Order(
		"DATE(reviewed_at)",
	).Scan(
		&results,
	).Error

	if err != nil {
		return nil, err
	}

	stats := make(
		[]analyticsDTO.DailyStat,
		0,
		len(results),
	)

	for _, result := range results {

		stats = append(
			stats,
			analyticsDTO.DailyStat{
				Date:  result.Date,
				Count: result.Count,
			},
		)
	}

	return stats, nil
}

func (r *analyticsRepository) GetStudySessionsPerDay(
	ctx context.Context,
	userID uuid.UUID,
	days int,
) ([]analyticsDTO.DailyStat, error) {

	var results []dailyStatResult

	startDate := time.Now().AddDate(
		0,
		0,
		-days,
	)

	err := r.db.WithContext(
		ctx,
	).Table(
		"study_sessions",
	).Select(
		"DATE(session_date) as date, COUNT(*) as count",
	).Where(
		"user_id = ?",
		userID,
	).Where(
		"session_date >= ?",
		startDate,
	).Group(
		"DATE(session_date)",
	).Order(
		"DATE(session_date)",
	).Scan(
		&results,
	).Error

	if err != nil {
		return nil, err
	}

	stats := make(
		[]analyticsDTO.DailyStat,
		0,
		len(results),
	)

	for _, result := range results {

		stats = append(
			stats,
			analyticsDTO.DailyStat{
				Date:  result.Date,
				Count: result.Count,
			},
		)
	}

	return stats, nil
}

func (r *analyticsRepository) GetVocabularyGrowth(
	ctx context.Context,
	userID uuid.UUID,
	days int,
) ([]analyticsDTO.DailyStat, error) {

	var results []dailyStatResult

	startDate := time.Now().AddDate(
		0,
		0,
		-days,
	)

	err := r.db.WithContext(
		ctx,
	).Table(
		"vocabularies",
	).Select(
		"DATE(created_at) as date, COUNT(*) as count",
	).Where(
		"user_id = ?",
		userID,
	).Where(
		"created_at >= ?",
		startDate,
	).Group(
		"DATE(created_at)",
	).Order(
		"DATE(created_at)",
	).Scan(
		&results,
	).Error

	if err != nil {
		return nil, err
	}

	stats := make(
		[]analyticsDTO.DailyStat,
		0,
		len(results),
	)

	for _, result := range results {

		stats = append(
			stats,
			analyticsDTO.DailyStat{
				Date:  result.Date,
				Count: result.Count,
			},
		)
	}

	return stats, nil
}

func (r *analyticsRepository) GetGoalCompletionRate(
	ctx context.Context,
	userID uuid.UUID,
) (float64, error) {

	var total int64
	var completed int64

	err := r.db.WithContext(
		ctx,
	).Table(
		"goals",
	).Where(
		"user_id = ?",
		userID,
	).Count(
		&total,
	).Error

	if err != nil {
		return 0, err
	}

	if total == 0 {
		return 0, nil
	}

	err = r.db.WithContext(
		ctx,
	).Table(
		"goals",
	).Where(
		"user_id = ?",
		userID,
	).Where(
		"status = ?",
		"COMPLETED",
	).Count(
		&completed,
	).Error

	if err != nil {
		return 0, err
	}

	return (float64(completed) / float64(total)) * 100, nil
}