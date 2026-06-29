package repositories

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type grammarRepository struct {
	db *gorm.DB
}

func NewGrammarRepository(
	db *gorm.DB,
) repositoryInterfaces.GrammarRepository {

	return &grammarRepository{
		db: db,
	}
}

func (r *grammarRepository) Create(
	ctx context.Context,
	grammar *models.Grammar,
) error {

	return r.db.WithContext(
		ctx,
	).Create(grammar).Error
}

func (r *grammarRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Grammar, error) {

	var grammar models.Grammar

	err := r.db.WithContext(ctx).
		First(&grammar, "id = ?", id).
		Error

	if err != nil {
		return nil, err
	}

	return &grammar, nil
}

func (r *grammarRepository) FindByUserAndPattern(
	ctx context.Context,
	userID uuid.UUID,
	pattern string,
) (*models.Grammar, error) {

	var grammar models.Grammar

	err := r.db.WithContext(ctx).
		Where(
			"user_id = ? AND pattern = ?",
			userID,
			pattern,
		).
		First(&grammar).Error

	if err != nil {
		return nil, err
	}

	return &grammar, nil
}

func (r *grammarRepository) FindAllByUserID(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.Grammar, error) {

	var grammars []models.Grammar

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&grammars).Error

	if err != nil {
		return nil, err
	}

	return grammars, nil
}

func (r *grammarRepository) FindFiltered(
	ctx context.Context,
	userID uuid.UUID,
	filter models.ListFilter,
) ([]models.Grammar, error) {

	var grammars []models.Grammar

	query := r.db.WithContext(ctx).
		Where("user_id = ?", userID)

	if filter.Search != "" {
		query = query.Where("pattern ILIKE ?", "%"+filter.Search+"%")
	}

	if filter.Favourite != nil {
		query = query.Where("favourite = ?", *filter.Favourite)
	}

	if filter.TagID != nil {
		query = query.Joins(
			"JOIN taggables ON taggables.item_id = grammars.id AND taggables.deleted_at IS NULL",
		).Where(
			"taggables.tag_id = ? AND taggables.item_type = 'GRAMMAR'",
			*filter.TagID,
		)
	}

	switch filter.SortBy {
	case "oldest":
		query = query.Order("created_at ASC")
	default:
		query = query.Order("created_at DESC")
	}

	err := query.Find(&grammars).Error

	if err != nil {
		return nil, err
	}

	return grammars, nil
}

func (r *grammarRepository) Update(
	ctx context.Context,
	grammar *models.Grammar,
) error {

	return r.db.WithContext(
		ctx,
	).Save(grammar).Error
}

func (r *grammarRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {

	return r.db.WithContext(ctx).
		Unscoped().
		Delete(&models.Grammar{}, id).
		Error
}

func (r *grammarRepository) FindRandomByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*models.Grammar, error) {

	var grammar models.Grammar

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("RANDOM()").
		First(&grammar).
		Error

	if err != nil {
		return nil, err
	}

	return &grammar, nil
}