package services

import (
	"context"
	"errors"

	grammarDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/grammar"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterface "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type grammarService struct {
	grammarRepository repositoryInterfaces.GrammarRepository
}

func NewGrammarService(
	grammarRepository repositoryInterfaces.GrammarRepository,
) serviceInterface.GrammarService {

	return &grammarService{
		grammarRepository: grammarRepository,
	}
}

func (s *grammarService) mapGrammarResponse(
	grammar *models.Grammar,
) *grammarDTO.GrammarResponse {

	return &grammarDTO.GrammarResponse{
		ID: grammar.ID,
		Pattern: grammar.Pattern,
		Meaning: grammar.Meaning,
		Example: grammar.Example,
		Note: grammar.Note,
		Favourite: grammar.Favourite,
	}
}

func (s *grammarService) CreateGrammar(
	ctx context.Context,
	userID uuid.UUID,
	dto grammarDTO.CreateGrammarRequest,
) (*grammarDTO.GrammarResponse, error) {

	existingGrammar, err :=
		s.grammarRepository.FindByUserAndPattern(
			ctx,
			userID,
			dto.Pattern,
		)

	if err != nil &&
		!errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingGrammar != nil {
		return nil, customErrors.ErrGrammarAlreadyExists
	}

	grammar := models.Grammar{
		UserID: userID,
		Pattern: dto.Pattern,
		Meaning: dto.Meaning,
		Example: dto.Example,
		Note: dto.Note,
	}

	if err := s.grammarRepository.Create(
		ctx,
		&grammar,
	); err != nil {
		return nil, err
	}

	return s.mapGrammarResponse(
		&grammar,
	), nil
}

func (s *grammarService) GetGrammarByID(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
) (*grammarDTO.GrammarResponse, error) {

	grammar, err := s.grammarRepository.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return nil, customErrors.ErrGrammarNotFound
	}

	if grammar.UserID != userID {
		return nil, customErrors.ErrGrammarNotFound
	}

	return s.mapGrammarResponse(
		grammar,
	), nil
}

func (s *grammarService) GetGrammarList(
	ctx context.Context,
	userID uuid.UUID,
	search string,
) ([]grammarDTO.GrammarResponse, error) {

	var (
		grammars []models.Grammar
		err error
	)

	if search != "" {

		grammars, err =
			s.grammarRepository.SearchByUserID(
				ctx,
				userID,
				search,
			)

	} else {

		grammars, err =
			s.grammarRepository.FindAllByUserID(
				ctx,
				userID,
			)
	}

	if err != nil {
		return nil, err
	}

	response :=
		make(
			[]grammarDTO.GrammarResponse,
			0,
		)

	for _, grammar := range grammars {

		response = append(
			response,
			*s.mapGrammarResponse(
				&grammar,
			),
		)
	}

	return response, nil
}

func (s *grammarService) UpdateGrammar(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
	dto grammarDTO.UpdateGrammarRequest,
) (*grammarDTO.GrammarResponse, error) {

	grammar, err := s.grammarRepository.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return nil, customErrors.ErrGrammarNotFound
	}

	if grammar.UserID != userID {
		return nil, customErrors.ErrGrammarNotFound
	}

	grammar.Pattern = dto.Pattern
	grammar.Meaning = dto.Meaning
	grammar.Example = dto.Example
	grammar.Note = dto.Note

	err = s.grammarRepository.Update(
		ctx,
		grammar,
	)

	if err != nil {
		return nil, err
	}

	return s.mapGrammarResponse(
		grammar,
	), nil
}

func (s *grammarService) DeleteGrammar(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
) error {

	grammar, err := s.grammarRepository.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return customErrors.ErrGrammarNotFound
	}

	if grammar.UserID != userID {
		return customErrors.ErrGrammarNotFound
	}

	return s.grammarRepository.Delete(
		ctx,
		id,
	)
}

func (s *grammarService) ToggleFavourite(
	ctx context.Context,
	userID uuid.UUID,
	id uuid.UUID,
) (*grammarDTO.GrammarResponse, error) {

	grammar, err := s.grammarRepository.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return nil, customErrors.ErrGrammarNotFound
	}

	if grammar.UserID != userID {
		return nil, customErrors.ErrGrammarNotFound
	}

	grammar.Favourite = !grammar.Favourite

	err = s.grammarRepository.Update(
		ctx,
		grammar,
	)

	if err != nil {
		return nil, err
	}

	return s.mapGrammarResponse(
		grammar,
	), nil
}

