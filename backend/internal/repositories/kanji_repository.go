package repositories

import (
	"context"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type kanjiRepository struct {
	db *gorm.DB
}

func NewKanjiRepository(
	db *gorm.DB,
) repositoryInterfaces.KanjiRepository {

	return &kanjiRepository{
		db: db,
	}
}

func (r *kanjiRepository) Create(
	ctx context.Context,
	kanji *models.Kanji,
) error {

	return r.db.WithContext(
		ctx,
	).Create(kanji).Error
}

func (r *kanjiRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Kanji, error) {

	var kanji models.Kanji

	err := r.db.WithContext(ctx).
		First(&kanji, "id = ?", id).
		Error

	if err != nil {
		return nil, err
	}

	return &kanji, nil
}

func (r *kanjiRepository) FindByUserAndCharacter(
	ctx context.Context,
	userID uuid.UUID,
	character string,
) (*models.Kanji, error) {

	var kanji models.Kanji

	err := r.db.WithContext(ctx).
		Where(
			"user_id = ? AND character = ?",
			userID,
			character,
		).
		First(&kanji).Error

	if err != nil {
		return nil, err
	}

	return &kanji, nil
}

func (r *kanjiRepository) FindAllByUserID(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.Kanji, error) {

	var kanjis []models.Kanji

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&kanjis).Error

	if err != nil {
		return nil, err
	}

	return kanjis, nil
}

func (r *kanjiRepository) SearchByUserID(
	ctx context.Context,
	userID uuid.UUID,
	search string,
) ([]models.Kanji, error) {

	var kanjis []models.Kanji

	err := r.db.WithContext(ctx).
		Where(
			"user_id = ? AND character ILIKE ?",
			userID,
			"%"+search+"%",
		).
		Find(&kanjis).Error

	if err != nil {
		return nil, err
	}

	return kanjis, nil
}

func (r *kanjiRepository) Update(
	ctx context.Context,
	kanji *models.Kanji,
) error {

	return r.db.WithContext(
		ctx,
	).Save(kanji).Error
}

func (r *kanjiRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {

	return r.db.WithContext(ctx).
		Delete(&models.Kanji{}, id).
		Error
}

func (r *kanjiRepository) FindRandomByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*models.Kanji, error) {

	var kanji models.Kanji

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("RANDOM()").
		First(&kanji).
		Error

	if err != nil {
		return nil, err
	}

	return &kanji, nil
}