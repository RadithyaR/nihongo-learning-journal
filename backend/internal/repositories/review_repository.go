package repositories

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(
	db *gorm.DB,
) repositoryInterfaces.ReviewRepository {
	return &reviewRepository{
		db: db,
	}
}

func (r *reviewRepository) Create(
	ctx context.Context,
	review *models.ReviewLog,
) error {

	return r.db.WithContext(
		ctx,
	).Create(review).Error
}

func (r *reviewRepository) GetReviewCountByUser(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {

	var count int64

	err := r.db.WithContext(ctx).
		Model(&models.ReviewLog{}).
		Where("user_id = ?", userID).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}