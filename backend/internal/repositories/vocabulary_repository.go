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
	).Unscoped().Delete(
		&models.Vocabulary{},
		"id = ?",
		id,
	).Error
}

func (r *vocabularyRepository) DeleteMeaningsByVocabularyID(
	ctx context.Context,
	vocabularyID uuid.UUID,
) error {

	return r.db.WithContext(
		ctx,
	).Where(
		"vocabulary_id = ?",
		vocabularyID,
	).Delete(
		&models.VocabularyMeaning{},
	).Error
}

func (r *vocabularyRepository) CreateMeanings(
	ctx context.Context,
	meanings []models.VocabularyMeaning,
) error {

	return r.db.WithContext(
		ctx,
	).Create(
		meanings,
	).Error
}

func (r *vocabularyRepository) FindFiltered(
	ctx context.Context,
	userID uuid.UUID,
	filter models.ListFilter,
) ([]models.Vocabulary, error) {

	var vocabularies []models.Vocabulary

	query := r.db.WithContext(ctx).
		Preload("Meanings").
		Where("user_id = ?", userID)

	if filter.Search != "" {
		query = query.Where("word ILIKE ?", "%"+filter.Search+"%")
	}

	if filter.Favourite != nil {
		query = query.Where("favourite = ?", *filter.Favourite)
	}

	if filter.TagID != nil {
		query = query.Joins(
			"JOIN taggables ON taggables.item_id = vocabularies.id AND taggables.deleted_at IS NULL",
		).Where(
			"taggables.tag_id = ? AND taggables.item_type = 'VOCABULARY'",
			*filter.TagID,
		)
	}

	switch filter.SortBy {
	case "oldest":
		query = query.Order("created_at ASC")
	default:
		query = query.Order("created_at DESC")
	}

	err := query.Find(&vocabularies).Error

	if err != nil {
		return nil, err
	}

	return vocabularies, nil
}

func (r *vocabularyRepository) FindRandomByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*models.Vocabulary, error) {

	var vocabulary models.Vocabulary

	err := r.db.WithContext(ctx).
		Preload("Meanings").
		Where(
			"user_id = ?",
			userID,
		).
		Order("RANDOM()").
		First(&vocabulary).Error

	if err != nil {
		return nil, err
	}

	return &vocabulary, nil
}