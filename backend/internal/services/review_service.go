package services

import (
	"context"
	"time"

	reviewDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/review"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/google/uuid"
)

type reviewService struct {
	reviewRepository repositoryInterfaces.ReviewRepository
	vocabularyRepository repositoryInterfaces.VocabularyRepository
}

func NewReviewService(
	reviewRepository repositoryInterfaces.ReviewRepository,
	vocabularyRepository repositoryInterfaces.VocabularyRepository,
) repositoryInterfaces.ReviewService {

	return &reviewService{
		reviewRepository: reviewRepository,
		vocabularyRepository: vocabularyRepository,
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
		VocabularyID: dto.VocabularyID,
		Rating: dto.Rating,
		ReviewedAt: time.Now(),
	}

	vocabulary, err :=
	s.vocabularyRepository.FindByID(
		ctx,
		dto.VocabularyID,
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