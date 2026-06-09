package repositories

import (
	"context"
	"time"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type studySessionRepository struct {
	db *gorm.DB
}

func NewStudySessionRepository(
	db *gorm.DB,
) repositoryInterfaces.StudySessionRepository {
	return &studySessionRepository{
		db: db,
	}
}

func (r *studySessionRepository) Create(
	ctx context.Context,
	session *models.StudySession,
) error {

	return r.db.WithContext(
		ctx,
	).Create(
		session,
	).Error
}

func (r *studySessionRepository) FindByUserAndDate(
	ctx context.Context,
	userID uuid.UUID,
	sessionDate time.Time,
) (*models.StudySession, error) {

	var session models.StudySession

	err := r.db.WithContext(
		ctx,
	).Where(
		"user_id = ?",
		userID,
	).Where(
		"session_date = ?",
		sessionDate.Format("2006-01-02"),
	).First(
		&session,
	).Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *studySessionRepository) Update(
	ctx context.Context,
	session *models.StudySession,
) error {

	return r.db.WithContext(
		ctx,
	).Save(
		session,
	).Error
}

func (r *studySessionRepository) AddItem(
	ctx context.Context,
	item *models.StudySessionItem,
) error {

	return r.db.WithContext(
		ctx,
	).Create(
		item,
	).Error
}

func (r *studySessionRepository) ItemExists(
	ctx context.Context,
	studySessionID uuid.UUID,
	itemType string,
	itemID uuid.UUID,
) (bool, error) {

	var count int64

	err := r.db.WithContext(
		ctx,
	).Model(
		&models.StudySessionItem{},
	).Where(
		"study_session_id = ?",
		studySessionID,
	).Where(
		"item_type = ?",
		itemType,
	).Where(
		"item_id = ?",
		itemID,
	).Count(
		&count,
	).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *studySessionRepository) FindHistoryByUserID(
	ctx context.Context,
	userID uuid.UUID,
	page int,
	limit int,
) ([]models.StudySession, int64, error) {

	var sessions []models.StudySession
	var total int64

	offset := (page - 1) * limit

	query := r.db.WithContext(
		ctx,
	).Model(
		&models.StudySession{},
	).Where(
		"user_id = ?",
		userID,
	)

	if err := query.Count(
		&total,
	).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order(
		"session_date DESC",
	).Offset(
		offset,
	).Limit(
		limit,
	).Find(
		&sessions,
	).Error

	if err != nil {
		return nil, 0, err
	}

	return sessions, total, nil
}

func (r *studySessionRepository) FindItemsBySessionID(
	ctx context.Context,
	sessionID uuid.UUID,
) ([]models.StudySessionItem, error) {

	var items []models.StudySessionItem

	err := r.db.WithContext(
		ctx,
	).Where(
		"study_session_id = ?",
		sessionID,
	).Find(
		&items,
	).Error

	if err != nil {
		return nil, err
	}

	return items, nil
}