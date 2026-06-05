package repositories

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type passwordResetRepository struct {
	db *gorm.DB
}

func NewPasswordResetRepository(
	db *gorm.DB,
) repositoryInterfaces.PasswordResetRepository {
	return &passwordResetRepository{
		db: db,
	}
}

func (r *passwordResetRepository) Create(
	ctx context.Context,
	reset *models.PasswordReset,
) error {

	return r.db.
		WithContext(ctx).
		Create(reset).
		Error
}

func (r *passwordResetRepository) FindByToken(
	ctx context.Context,
	token string,
) (*models.PasswordReset, error) {

	var reset models.PasswordReset

	err := r.db.
		WithContext(ctx).
		Where("token = ?", token).
		First(&reset).
		Error

	if err != nil {
		return nil, err
	}

	return &reset, nil
}

func (r *passwordResetRepository) FindLatestByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*models.PasswordReset, error) {

	var reset models.PasswordReset

	err := r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		First(&reset).
		Error

	if err != nil {
		return nil, err
	}

	return &reset, nil
}

func (r *passwordResetRepository) Update(
	ctx context.Context,
	reset *models.PasswordReset,
) error {

	return r.db.
		WithContext(ctx).
		Save(reset).
		Error
}

func (r *passwordResetRepository) DeleteByUserID(
	ctx context.Context,
	userID uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&models.PasswordReset{}).
		Error
}