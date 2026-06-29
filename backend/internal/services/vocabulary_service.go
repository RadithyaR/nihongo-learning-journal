package services

import (
	"context"

	vocabularyDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/vocabulary"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterface "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type vocabularyService struct {
	vocabularyRepository repositoryInterfaces.VocabularyRepository
}

func NewVocabularyService(
	vocabularyRepository repositoryInterfaces.VocabularyRepository,
) serviceInterface.VocabularyService {
	return &vocabularyService{
		vocabularyRepository: vocabularyRepository,
	}
}

func (s *vocabularyService) CreateVocabulary(
	ctx context.Context,
	userID uuid.UUID,
	dto vocabularyDTO.CreateVocabularyRequest,
) (*vocabularyDTO.VocabularyResponse, error) {

	existingVocabulary, err :=
		s.vocabularyRepository.FindByUserAndWord(
			ctx,
			userID,
			dto.Word,
		)

	if err != nil &&
		err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existingVocabulary != nil {
		return nil,
			customErrors.ErrVocabularyAlreadyExists
	}

	vocabulary := models.Vocabulary{
		UserID:    userID,
		Word:      dto.Word,
		Reading:   dto.Reading,
		Source:    dto.Source,
		Note:      dto.Note,
		Favourite: false,
	}

	for index, meaning := range dto.Meanings {
		vocabulary.Meanings =
			append(
				vocabulary.Meanings,
				models.VocabularyMeaning{
					Meaning: meaning,
					OrderNumber: index + 1,
				},
			)
	}

	err = s.vocabularyRepository.Create(
		ctx,
		&vocabulary,
	)

	if err != nil {
		return nil, err
	}

	return s.mapVocabularyResponse(
		&vocabulary,
	), nil
}

func (s *vocabularyService) mapVocabularyResponse(
	vocabulary *models.Vocabulary,
) *vocabularyDTO.VocabularyResponse {

	meanings :=
		make(
			[]vocabularyDTO.VocabularyMeaningResponse,
			0,
		)

	for _, meaning := range vocabulary.Meanings {

		meanings =
			append(
				meanings,
				vocabularyDTO.VocabularyMeaningResponse{
					ID: meaning.ID,
					Meaning: meaning.Meaning,
					OrderNumber: meaning.OrderNumber,
				},
			)
	}

	return &vocabularyDTO.VocabularyResponse{
		ID: vocabulary.ID,

		Word: vocabulary.Word,

		Reading: vocabulary.Reading,

		Source: vocabulary.Source,

		Note: vocabulary.Note,

		Status: vocabulary.Status,

		Favourite: vocabulary.Favourite,

		Meanings: meanings,
	}
}

func (s *vocabularyService) GetVocabularyByID(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
) (*vocabularyDTO.VocabularyResponse, error) {

	vocabulary, err :=
		s.vocabularyRepository.FindByID(
			ctx,
			id,
		)

	if err != nil {
		return nil,
			customErrors.ErrVocabularyNotFound
	}

	if vocabulary.UserID != userID {
		return nil,
			customErrors.ErrVocabularyNotFound
	}

	return s.mapVocabularyResponse(
		vocabulary,
	), nil
}

func (s *vocabularyService) GetVocabularyList(
	ctx context.Context,
	userID uuid.UUID,
	filter models.ListFilter,
) ([]vocabularyDTO.VocabularyResponse, error) {

	vocabularies, err := s.vocabularyRepository.FindFiltered(
		ctx,
		userID,
		filter,
	)

	if err != nil {
		return nil, err
	}

	response :=
		make(
			[]vocabularyDTO.VocabularyResponse,
			0,
		)

	for _, vocabulary := range vocabularies {

		response =
			append(
				response,
				*s.mapVocabularyResponse(
					&vocabulary,
				),
			)
	}

	return response, nil
}

func (s *vocabularyService) UpdateVocabulary(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
	dto vocabularyDTO.UpdateVocabularyRequest,
) (*vocabularyDTO.VocabularyResponse, error) {

	vocabulary, err :=
		s.vocabularyRepository.FindByID(
			ctx,
			id,
		)

	if err != nil {
		return nil,
			customErrors.ErrVocabularyNotFound
	}

	if vocabulary.UserID != userID {
		return nil,
			customErrors.ErrVocabularyNotFound
	}

	vocabulary.Word = dto.Word
	vocabulary.Reading = dto.Reading
	vocabulary.Source = dto.Source
	vocabulary.Note = dto.Note

	err = s.vocabularyRepository.Update(
		ctx,
		vocabulary,
	)

	if err != nil {
		return nil, err
	}

	err = s.vocabularyRepository.DeleteMeaningsByVocabularyID(
		ctx,
		vocabulary.ID,
	)

	if err != nil {
		return nil, err
	}

	var newMeanings []models.VocabularyMeaning
	for i, meaning := range dto.Meanings {
		newMeanings = append(
			newMeanings,
			models.VocabularyMeaning{
				VocabularyID: vocabulary.ID,
				Meaning:      meaning,
				OrderNumber:  i + 1,
			},
		)
	}

	err = s.vocabularyRepository.CreateMeanings(
		ctx,
		newMeanings,
	)

	if err != nil {
		return nil, err
	}

	vocabulary.Meanings = newMeanings

	return s.mapVocabularyResponse(
		vocabulary,
	), nil
}
func (s *vocabularyService) DeleteVocabulary(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
) error {

	vocabulary, err :=
		s.vocabularyRepository.FindByID(
			ctx,
			id,
		)

	if err != nil {
		return customErrors.ErrVocabularyNotFound
	}

	if vocabulary.UserID != userID {
		return customErrors.ErrVocabularyNotFound
	}

	return s.vocabularyRepository.Delete(
		ctx,
		id,
	)
}

func (s *vocabularyService) ToggleFavourite(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
) (*vocabularyDTO.VocabularyResponse, error) {

	vocabulary, err :=
		s.vocabularyRepository.FindByID(
			ctx,
			id,
		)

	if err != nil {
		return nil,
			customErrors.ErrVocabularyNotFound
	}

	if vocabulary.UserID != userID {
		return nil,
			customErrors.ErrVocabularyNotFound
	}

	vocabulary.Favourite =
		!vocabulary.Favourite

	err =
		s.vocabularyRepository.Update(
			ctx,
			vocabulary,
		)

	if err != nil {
		return nil, err
	}

	return s.mapVocabularyResponse(
		vocabulary,
	), nil
}
