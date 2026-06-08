package services

import (
	"context"
	"time"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/constants"
	reviewDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/review"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/google/uuid"
)

type reviewService struct {
	reviewRepository repositoryInterfaces.ReviewRepository
	vocabularyRepository repositoryInterfaces.VocabularyRepository
	kanjiRepository repositoryInterfaces.KanjiRepository
}

func NewReviewService(
	reviewRepository repositoryInterfaces.ReviewRepository,
	vocabularyRepository repositoryInterfaces.VocabularyRepository,
	kanjiRepository repositoryInterfaces.KanjiRepository,
) repositoryInterfaces.ReviewService {

	return &reviewService{
		reviewRepository: reviewRepository,
		vocabularyRepository: vocabularyRepository,
		kanjiRepository: kanjiRepository,
	}
}

func (s *reviewService) GetNextReview(
	ctx context.Context,
	userID uuid.UUID,
) (*reviewDTO.NextReviewResponse, error) {

	vocabulary, err :=
		s.vocabularyRepository.FindRandomByUserID(
			ctx,
			userID,
		)

	if err != nil {
		return nil, err
	}

	return &reviewDTO.NextReviewResponse{
		ID: vocabulary.ID,
		Word: vocabulary.Word,
		Reading: vocabulary.Reading,
	}, nil
}

func (s *reviewService) SubmitReview(
	ctx context.Context,
	userID uuid.UUID,
	dto reviewDTO.SubmitReviewRequest,
) error {

	review := models.ReviewLog{
	UserID: userID,
	ItemType: constants.ReviewTypeVocabulary,
	ItemID: dto.ItemID,
	Rating: dto.Rating,
	ReviewedAt: time.Now(),
}

	vocabulary, err :=
	s.vocabularyRepository.FindByID(
		ctx,
		dto.ItemID,
	)

	if err != nil {
		return customErrors.ErrVocabularyNotFound
	}

	if vocabulary.UserID != userID {
		return customErrors.ErrVocabularyNotFound
	}

	return s.reviewRepository.Create(
		ctx,
		&review,
	)
}

func (s *reviewService) GetNextKanjiReview(
	ctx context.Context,
	userID uuid.UUID,
) (*reviewDTO.NextKanjiReviewResponse, error) {

	kanji, err :=
		s.kanjiRepository.FindRandomByUserID(
			ctx,
			userID,
		)

	if err != nil {
		return nil, err
	}

	return &reviewDTO.NextKanjiReviewResponse{
		ID:        kanji.ID,
		Character: kanji.Character,
		Meaning:   kanji.Meaning,
		Onyomi:    kanji.Onyomi,
		Kunyomi:   kanji.Kunyomi,
	}, nil
}

func (s *reviewService) SubmitKanjiReview(
	ctx context.Context,
	userID uuid.UUID,
	dto reviewDTO.SubmitReviewRequest,
) error {

	kanji, err :=
		s.kanjiRepository.FindByID(
			ctx,
			dto.ItemID,
		)

	if err != nil {
		return err
	}

	if kanji.UserID != userID {
		return customErrors.ErrKanjiNotFound
	}

	review := models.ReviewLog{
		UserID:     userID,
		ItemType:   constants.ReviewTypeKanji,
		ItemID:     dto.ItemID,
		Rating:     dto.Rating,
		ReviewedAt: time.Now(),
	}

	return s.reviewRepository.Create(
		ctx,
		&review,
	)
}