package repositories

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type taggableRepository struct {
	db *gorm.DB
}

func NewTaggableRepository(
	db *gorm.DB,
) repositoryInterfaces.TaggableRepository {
	return &taggableRepository{
		db: db,
	}
}

func (r *taggableRepository) Create(
	ctx context.Context,
	taggable *models.Taggable,
) error {

	return r.db.WithContext(
		ctx,
	).Create(
		taggable,
	).Error
}

func (r *taggableRepository) Exists(
	ctx context.Context,
	tagID uuid.UUID,
	itemType string,
	itemID uuid.UUID,
) (bool, error) {

	var count int64

	err := r.db.WithContext(
		ctx,
	).Model(
		&models.Taggable{},
	).Where(
		"tag_id = ?",
		tagID,
	).Where(
		"item_type = ?",
		itemType,
	).Where(
		"item_id = ?",
		itemID,
	).Where(
		"deleted_at IS NULL",
	).Count(
		&count,
	).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *taggableRepository) Delete(
	ctx context.Context,
	tagID uuid.UUID,
	itemType string,
	itemID uuid.UUID,
) error {

	return r.db.WithContext(
		ctx,
	).Where(
		"tag_id = ?",
		tagID,
	).Where(
		"item_type = ?",
		itemType,
	).Where(
		"item_id = ?",
		itemID,
	).Unscoped().Delete(
		&models.Taggable{},
	).Error
}

func (r *taggableRepository) FindTagsByItem(
	ctx context.Context,
	itemType string,
	itemID uuid.UUID,
) ([]models.Tag, error) {

	var tags []models.Tag

	err := r.db.WithContext(
		ctx,
	).Model(
		&models.Tag{},
	).Joins(
		"JOIN taggables ON taggables.tag_id = tags.id AND taggables.deleted_at IS NULL",
	).Where(
		"taggables.item_type = ?",
		itemType,
	).Where(
		"taggables.item_id = ?",
		itemID,
	).Find(
		&tags,
	).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}