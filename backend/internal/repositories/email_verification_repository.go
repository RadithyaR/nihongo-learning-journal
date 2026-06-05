package repositories

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type emailVerificationRepository struct {
	db *gorm.DB
}

func NewEmailVerificationRepository(
	db *gorm.DB,
) repositoryInterfaces.EmailVerificationRepository {
	return &emailVerificationRepository{
		db: db,
	}
}

func (r *emailVerificationRepository) Create(
	ctx context.Context,
	verification *models.EmailVerification,
) error {

	return r.db.
		WithContext(ctx).
		Create(verification).
		Error
}

func (r *emailVerificationRepository) FindByToken(
	ctx context.Context,
	token string,
) (*models.EmailVerification, error) {

	var verification models.EmailVerification

	err := r.db.
		WithContext(ctx).
		Where("token = ?", token).
		First(&verification).
		Error

	if err != nil {
		return nil, err
	}

	return &verification, nil
}

func (r *emailVerificationRepository) DeleteByUserID(
	ctx context.Context,
	userID uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&models.EmailVerification{}).
		Error
}

func (r *emailVerificationRepository) Update(
	ctx context.Context,
	verification *models.EmailVerification,
) error {

	return r.db.
		WithContext(ctx).
		Save(verification).
		Error
}

func (r *emailVerificationRepository) FindLatestByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*models.EmailVerification, error) {

	var verification models.EmailVerification

	err := r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		First(&verification).
		Error

	if err != nil {
		return nil, err
	}

	return &verification, nil
}