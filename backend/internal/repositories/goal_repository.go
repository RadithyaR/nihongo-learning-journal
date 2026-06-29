package repositories

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type goalRepository struct {
	db *gorm.DB
}

func NewGoalRepository(
	db *gorm.DB,
) repositoryInterfaces.GoalRepository {
	return &goalRepository{
		db: db,
	}
}

func (r *goalRepository) Create(
	ctx context.Context,
	goal *models.Goal,
) error {

	return r.db.WithContext(
		ctx,
	).Create(
		goal,
	).Error
}

func (r *goalRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Goal, error) {

	var goal models.Goal

	err := r.db.WithContext(
		ctx,
	).First(
		&goal,
		"id = ?",
		id,
	).Error

	if err != nil {
		return nil, err
	}

	return &goal, nil
}

func (r *goalRepository) FindAllByUserID(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.Goal, error) {

	var goals []models.Goal

	err := r.db.WithContext(
		ctx,
	).Where(
		"user_id = ?",
		userID,
	).Order(
		"created_at DESC",
	).Find(
		&goals,
	).Error

	if err != nil {
		return nil, err
	}

	return goals, nil
}

func (r *goalRepository) Update(
	ctx context.Context,
	goal *models.Goal,
) error {

	return r.db.WithContext(
		ctx,
	).Save(
		goal,
	).Error
}

func (r *goalRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {

	return r.db.WithContext(
		ctx,
	).Unscoped().Delete(
		&models.Goal{},
		"id = ?",
		id,
	).Error
}

func (r *goalRepository) CountProgress(
	ctx context.Context,
	userID uuid.UUID,
	goalType string,
) (int64, error) {

	var count int64

	err := r.db.WithContext(
		ctx,
	).Table(
		"study_session_items ssi",
	).Joins(
		"JOIN study_sessions ss ON ss.id = ssi.study_session_id",
	).Where(
		"ss.user_id = ?",
		userID,
	).Where(
		"ssi.item_type = ?",
		goalType,
	).Distinct(
		"ssi.item_id",
	).Count(
		&count,
	).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}