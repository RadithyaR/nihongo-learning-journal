package repositories

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type userSessionRepository struct {
	db *gorm.DB
}

func NewUserSessionRepository(
	db *gorm.DB,
) repositoryInterfaces.UserSessionRepository {
	return &userSessionRepository{
		db: db,
	}
}

func (r *userSessionRepository) Create(
	ctx context.Context,
	session *models.UserSession,
) error {
	return r.db.
		WithContext(ctx).
		Create(session).
		Error
}
func (r *userSessionRepository) FindByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*models.UserSession, error) {

	var session models.UserSession

	err := r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		First(&session).
		Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *userSessionRepository) FindByTokenID(
	ctx context.Context,
	tokenID uuid.UUID,
) (*models.UserSession, error) {

	var session models.UserSession

	err := r.db.
		WithContext(ctx).
		Where("token_id = ?", tokenID).
		First(&session).
		Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *userSessionRepository) Update(
	ctx context.Context,
	session *models.UserSession,
) error {

	return r.db.
		WithContext(ctx).
		Save(session).
		Error
}
func (r *userSessionRepository) Delete(
	ctx context.Context,
	sessionID uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Delete(
			&models.UserSession{},
			"id = ?",
			sessionID,
		).
		Error
}

func (r *userSessionRepository) DeleteByTokenID(
	ctx context.Context,
	tokenID uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Delete(
			&models.UserSession{},
			"token_id = ?",
			tokenID,
		).
		Error
}