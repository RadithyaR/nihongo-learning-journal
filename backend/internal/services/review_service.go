package services

import (
	"context"
	"time"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/constants"
	reviewDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/review"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterface "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/google/uuid"
)

type reviewService struct {
	reviewRepository repositoryInterfaces.ReviewRepository
	vocabularyRepository repositoryInterfaces.VocabularyRepository
	kanjiRepository repositoryInterfaces.KanjiRepository
	grammarRepository repositoryInterfaces.GrammarRepository
	studySessionService serviceInterface.StudySessionService
}

func NewReviewService(
	reviewRepository repositoryInterfaces.ReviewRepository,
	vocabularyRepository repositoryInterfaces.VocabularyRepository,
	kanjiRepository repositoryInterfaces.KanjiRepository,
	grammarRepository repositoryInterfaces.GrammarRepository,
	studySessionService serviceInterface.StudySessionService,
) serviceInterface.ReviewService {

	return &reviewService{
		reviewRepository: reviewRepository,
		vocabularyRepository: vocabularyRepository,
		kanjiRepository: kanjiRepository,
		grammarRepository: grammarRepository,
		studySessionService: studySessionService,
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

	err = s.reviewRepository.Create(
		ctx,
		&review,
	)

	if err != nil {
		return err
	}

	return s.studySessionService.AddVocabulary(
		ctx,
		userID,
		dto.ItemID,
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

	err = s.reviewRepository.Create(
		ctx,
		&review,
	)

	if err != nil {
		return err
	}

	return s.studySessionService.AddKanji(
		ctx,
		userID,
		dto.ItemID,
)
}

func (s *reviewService) GetNextGrammarReview(
	ctx context.Context,
	userID uuid.UUID,
) (*reviewDTO.NextGrammarReviewResponse, error) {

	grammar, err :=
		s.grammarRepository.FindRandomByUserID(
			ctx,
			userID,
		)

	if err != nil {
		return nil, err
	}

	return &reviewDTO.NextGrammarReviewResponse{
		ID: grammar.ID,
		Pattern: grammar.Pattern,
		Meaning: grammar.Meaning,
	}, nil
}

func (s *reviewService) SubmitGrammarReview(
	ctx context.Context,
	userID uuid.UUID,
	dto reviewDTO.SubmitGrammarReviewRequest,
) error {
	grammar, err :=
		s.grammarRepository.FindByID(
			ctx,
			dto.GrammarID,
		)

	if err != nil {
		return customErrors.ErrGrammarNotFound
	}

	if grammar.UserID != userID {
		return customErrors.ErrGrammarNotFound
	}

	review := models.ReviewLog{
		UserID: userID,
		ItemType: constants.ReviewTypeGrammar,
		ItemID: dto.GrammarID,
		Rating: dto.Rating,
		ReviewedAt: time.Now(),
	}

	err = s.reviewRepository.Create(
		ctx,
		&review,
	)

	if err != nil {
		return err
	}

	return s.studySessionService.AddGrammar(
		ctx,
		userID,
		dto.GrammarID,
	)
}