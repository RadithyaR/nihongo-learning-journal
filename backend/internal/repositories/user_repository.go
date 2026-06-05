package repositories

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) repositoryInterfaces.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(
	ctx context.Context,
	user *models.User,
) error {
	return r.db.WithContext(ctx).
		Create(user).
		Error
}

func (r *userRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByGoogleID(
	ctx context.Context,
	googleID string,
) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Where("google_id = ?", googleID).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(
	ctx context.Context,
	user *models.User,
) error {

	return r.db.
		WithContext(ctx).
		Save(user).
		Error
}