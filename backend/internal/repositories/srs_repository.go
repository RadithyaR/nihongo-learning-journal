package repositories

import (
	"context"
	"time"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type srsRepository struct {
	db *gorm.DB
}

func NewSRSRepository(
	db *gorm.DB,
) repositoryInterfaces.SRSRepository {

	return &srsRepository{
		db: db,
	}
}

func (r *srsRepository) Create(
	ctx context.Context,
	record *models.SRSRecord,
) error {

	return r.db.WithContext(
		ctx,
	).Create(
		record,
	).Error
}

func (r *srsRepository) Update(
	ctx context.Context,
	record *models.SRSRecord,
) error {

	return r.db.WithContext(
		ctx,
	).Save(
		record,
	).Error
}

func (r *srsRepository) FindByItem(
	ctx context.Context,
	userID uuid.UUID,
	itemType string,
	itemID uuid.UUID,
) (*models.SRSRecord, error) {

	var record models.SRSRecord

	err := r.db.WithContext(
		ctx,
	).Where(
		"user_id = ?",
		userID,
	).Where(
		"item_type = ?",
		itemType,
	).Where(
		"item_id = ?",
		itemID,
	).First(
		&record,
	).Error

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *srsRepository) GetDueItems(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.SRSRecord, error) {

	var records []models.SRSRecord

	err := r.db.WithContext(
		ctx,
	).Where(
		"user_id = ?",
		userID,
	).Where(
		"next_review_at <= ?",
		time.Now(),
	).Order(
		"next_review_at ASC",
	).Find(
		&records,
	).Error

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (r *srsRepository) CountDueToday(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {

	var count int64

	now := time.Now()

	startOfDay := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		now.Location(),
	)

	endOfDay := startOfDay.Add(
		24*time.Hour,
	)

	err := r.db.WithContext(
		ctx,
	).Model(
		&models.SRSRecord{},
	).Where(
		"user_id = ?",
		userID,
	).Where(
		"next_review_at >= ?",
		startOfDay,
	).Where(
		"next_review_at < ?",
		endOfDay,
	).Count(
		&count,
	).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *srsRepository) CountOverdue(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {

	var count int64

	now := time.Now()

	err := r.db.WithContext(
		ctx,
	).Model(
		&models.SRSRecord{},
	).Where(
		"user_id = ?",
		userID,
	).Where(
		"next_review_at < ?",
		now,
	).Count(
		&count,
	).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *srsRepository) GetNextDueItem(
	ctx context.Context,
	userID uuid.UUID,
	itemType string,
) (*models.SRSRecord, error) {

	var record models.SRSRecord

	err := r.db.WithContext(
		ctx,
	).Where(
		"user_id = ?",
		userID,
	).Where(
		"item_type = ?",
		itemType,
	).Where(
		"next_review_at <= ?",
		time.Now(),
	).Order(
		"next_review_at ASC",
	).First(
		&record,
	).Error

	if err != nil {
		return nil, err
	}

	return &record, nil
}