package services

import (
	"context"
	"errors"

	kanjiDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/kanji"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterface "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type kanjiService struct {
	kanjiRepository repositoryInterfaces.KanjiRepository
}

func NewKanjiService(
	kanjiRepository repositoryInterfaces.KanjiRepository,
) serviceInterface.KanjiService {

	return &kanjiService{
		kanjiRepository: kanjiRepository,
	}
}

func (s *kanjiService) mapKanjiResponse(
	kanji *models.Kanji,
) *kanjiDTO.KanjiResponse {

	return &kanjiDTO.KanjiResponse{
		ID: kanji.ID,
		Character: kanji.Character,
		Meaning: kanji.Meaning,
		Onyomi: kanji.Onyomi,
		Kunyomi: kanji.Kunyomi,
		JLPTLevel: kanji.JLPTLevel,
		Favourite: kanji.Favourite,
	}
}

func (s *kanjiService) CreateKanji(
	ctx context.Context,
	userID uuid.UUID,
	dto kanjiDTO.CreateKanjiRequest,
) (*kanjiDTO.KanjiResponse, error) {

	existingKanji, err :=
		s.kanjiRepository.FindByUserAndCharacter(
			ctx,
			userID,
			dto.Character,
		)

	if err != nil &&
		!errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingKanji != nil {
		return nil, customErrors.ErrKanjiAlreadyExists
	}

	kanji := models.Kanji{
		UserID: userID,

		Character: dto.Character,

		Meaning: dto.Meaning,

		Onyomi: dto.Onyomi,

		Kunyomi: dto.Kunyomi,

		JLPTLevel: dto.JLPTLevel,
	}

	if err := s.kanjiRepository.Create(
		ctx,
		&kanji,
	); err != nil {
		return nil, err
	}

	return s.mapKanjiResponse(
		&kanji,
	), nil
}

func (s *kanjiService) GetKanjiByID(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
) (*kanjiDTO.KanjiResponse, error) {

	kanji, err := s.kanjiRepository.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return nil, customErrors.ErrKanjiNotFound
	}

	if kanji.UserID != userID {
		return nil, customErrors.ErrKanjiNotFound
	}

	return s.mapKanjiResponse(
		kanji,
	), nil
}

func (s *kanjiService) GetKanjiList(
	ctx context.Context,
	userID uuid.UUID,
	filter models.ListFilter,
) ([]kanjiDTO.KanjiResponse, error) {

	kanjis, err := s.kanjiRepository.FindFiltered(
		ctx,
		userID,
		filter,
	)

	if err != nil {
		return nil, err
	}

	response :=
		make(
			[]kanjiDTO.KanjiResponse,
			0,
		)

	for _, kanji := range kanjis {

		response =
			append(
				response,
				*s.mapKanjiResponse(
					&kanji,
				),
			)
	}

	return response, nil
}

func (s *kanjiService) UpdateKanji(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
	dto kanjiDTO.UpdateKanjiRequest,
) (*kanjiDTO.KanjiResponse, error) {

	kanji, err := s.kanjiRepository.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return nil, customErrors.ErrKanjiNotFound
	}

	if kanji.UserID != userID {
		return nil, customErrors.ErrKanjiNotFound
	}

	kanji.Character = dto.Character
	kanji.Meaning = dto.Meaning
	kanji.Onyomi = dto.Onyomi
	kanji.Kunyomi = dto.Kunyomi
	kanji.JLPTLevel = dto.JLPTLevel

	err = s.kanjiRepository.Update(
		ctx,
		kanji,
	)

	if err != nil {
		return nil, err
	}

	return s.mapKanjiResponse(
		kanji,
	), nil
}

func (s *kanjiService) DeleteKanji(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
) error {

	kanji, err := s.kanjiRepository.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return customErrors.ErrKanjiNotFound
	}

	if kanji.UserID != userID {
		return customErrors.ErrKanjiNotFound
	}

	return s.kanjiRepository.Delete(
		ctx,
		id,
	)
}

func (s *kanjiService) ToggleFavourite(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
) (*kanjiDTO.KanjiResponse, error) {

	kanji, err := s.kanjiRepository.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return nil, customErrors.ErrKanjiNotFound
	}

	if kanji.UserID != userID {
		return nil, customErrors.ErrKanjiNotFound
	}

	kanji.Favourite = !kanji.Favourite

	err = s.kanjiRepository.Update(
		ctx,
		kanji,
	)

	if err != nil {
		return nil, err
	}

	return s.mapKanjiResponse(
		kanji,
	), nil
}

