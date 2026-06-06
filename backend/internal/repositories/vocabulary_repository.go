package repositories

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type vocabularyRepository struct {
	db *gorm.DB
}

func NewVocabularyRepository(
	db *gorm.DB,
) repositoryInterfaces.VocabularyRepository {
	return &vocabularyRepository{
		db: db,
	}
}

func (r *vocabularyRepository) Create(
	ctx context.Context,
	vocabulary *models.Vocabulary,
) error {

	return r.db.WithContext(
		ctx,
	).Create(
		vocabulary,
	).Error
}

func (r *vocabularyRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Vocabulary, error) {

	var vocabulary models.Vocabulary

	err := r.db.WithContext(
		ctx,
	).
		Preload("Meanings").
		First(
			&vocabulary,
			"id = ?",
			id,
		).Error

	if err != nil {
		return nil, err
	}

	return &vocabulary, nil
}

func (r *vocabularyRepository) FindByUserAndWord(
	ctx context.Context,
	userID uuid.UUID,
	word string,
) (*models.Vocabulary, error) {

	var vocabulary models.Vocabulary

	err := r.db.WithContext(
		ctx,
	).
		Where(
			"user_id = ? AND word = ?",
			userID,
			word,
		).
		First(
			&vocabulary,
		).Error

	if err != nil {
		return nil, err
	}

	return &vocabulary, nil
}

func (r *vocabularyRepository) FindAllByUserID(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.Vocabulary, error) {

	var vocabularies []models.Vocabulary

	err := r.db.WithContext(
		ctx,
	).
		Preload("Meanings").
		Where(
			"user_id = ?",
			userID,
		).
		Find(
			&vocabularies,
		).Error

	if err != nil {
		return nil, err
	}

	return vocabularies, nil
}

func (r *vocabularyRepository) Update(
	ctx context.Context,
	vocabulary *models.Vocabulary,
) error {

	return r.db.WithContext(
		ctx,
	).Save(
		vocabulary,
	).Error
}

func (r *vocabularyRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {

	return r.db.WithContext(
		ctx,
	).Delete(
		&models.Vocabulary{},
		"id = ?",
		id,
	).Error
}