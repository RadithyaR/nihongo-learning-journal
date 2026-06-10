package repositories

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(
	db *gorm.DB,
) repositoryInterfaces.TagRepository {
	return &tagRepository{
		db: db,
	}
}

func (r *tagRepository) Create(
	ctx context.Context,
	tag *models.Tag,
) error {

	return r.db.WithContext(
		ctx,
	).Create(
		tag,
	).Error
}

func (r *tagRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Tag, error) {

	var tag models.Tag

	err := r.db.WithContext(
		ctx,
	).First(
		&tag,
		"id = ?",
		id,
	).Error

	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r *tagRepository) FindByUserID(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.Tag, error) {

	var tags []models.Tag

	err := r.db.WithContext(
		ctx,
	).Where(
		"user_id = ?",
		userID,
	).Order(
		"name ASC",
	).Find(
		&tags,
	).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *tagRepository) ExistsByName(
	ctx context.Context,
	userID uuid.UUID,
	name string,
) (bool, error) {

	var count int64

	err := r.db.WithContext(
		ctx,
	).Model(
		&models.Tag{},
	).Where(
		"user_id = ?",
		userID,
	).Where(
		"name = ?",
		name,
	).Count(
		&count,
	).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *tagRepository) Update(
	ctx context.Context,
	tag *models.Tag,
) error {

	return r.db.WithContext(
		ctx,
	).Save(
		tag,
	).Error
}

func (r *tagRepository) Delete(
	ctx context.Context,
	tag *models.Tag,
) error {

	return r.db.WithContext(
		ctx,
	).Delete(
		tag,
	).Error
}