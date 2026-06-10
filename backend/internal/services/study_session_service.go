package services

import (
	"context"
	"time"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/constants"
	studySessionDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/studySession"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterface "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/google/uuid"
)

type studySessionService struct {
	studySessionRepository repositoryInterfaces.StudySessionRepository
	vocabularyRepository   repositoryInterfaces.VocabularyRepository
	kanjiRepository        repositoryInterfaces.KanjiRepository
	grammarRepository      repositoryInterfaces.GrammarRepository
}

func NewStudySessionService(
	studySessionRepository repositoryInterfaces.StudySessionRepository,
	vocabularyRepository repositoryInterfaces.VocabularyRepository,
	kanjiRepository repositoryInterfaces.KanjiRepository,
	grammarRepository repositoryInterfaces.GrammarRepository,
) serviceInterface.StudySessionService {

	return &studySessionService{
		studySessionRepository: studySessionRepository,
		vocabularyRepository:   vocabularyRepository,
		kanjiRepository:        kanjiRepository,
		grammarRepository:      grammarRepository,
	}
}

func (s *studySessionService) getOrCreateTodaySession(
	ctx context.Context,
	userID uuid.UUID,
) (*models.StudySession, error) {

	session, err :=
		s.studySessionRepository.FindByUserAndDate(
			ctx,
			userID,
			time.Now(),
		)

	if err == nil {
		return session, nil
	}

	session = &models.StudySession{
		ID:          uuid.New(),
		UserID:      userID,
		SessionDate: time.Now(),
	}

	err = s.studySessionRepository.Create(
		ctx,
		session,
	)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *studySessionService) AddVocabulary(
	ctx context.Context,
	userID uuid.UUID,
	vocabularyID uuid.UUID,
) error {

	session, err :=
		s.getOrCreateTodaySession(
			ctx,
			userID,
		)

	if err != nil {
		return err
	}

	exists, err :=
		s.studySessionRepository.ItemExists(
			ctx,
			session.ID,
			constants.StudyItemVocabulary,
			vocabularyID,
		)

	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	item := models.StudySessionItem{
		ID:             uuid.New(),
		StudySessionID: session.ID,
		ItemType:       constants.StudyItemVocabulary,
		ItemID:         vocabularyID,
	}

	return s.studySessionRepository.AddItem(
		ctx,
		&item,
	)
}

func (s *studySessionService) AddKanji(
	ctx context.Context,
	userID uuid.UUID,
	kanjiID uuid.UUID,
) error {


	session, err :=
		s.getOrCreateTodaySession(
			ctx,
			userID,
		)

	if err != nil {
		return err
	}

	exists, err :=
		s.studySessionRepository.ItemExists(
			ctx,
			session.ID,
			constants.StudyItemVocabulary,
			kanjiID,
		)

	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	item := models.StudySessionItem{
		ID:             uuid.New(),
		StudySessionID: session.ID,
		ItemType:       constants.StudyItemKanji,
		ItemID:         kanjiID,
	}

	return s.studySessionRepository.AddItem(
		ctx,
		&item,
	)
}

func (s *studySessionService) AddGrammar(
	ctx context.Context,
	userID uuid.UUID,
	grammarID uuid.UUID,
) error {

	session, err :=
		s.getOrCreateTodaySession(
			ctx,
			userID,
		)

	if err != nil {
		return err
	}

	exists, err :=
		s.studySessionRepository.ItemExists(
			ctx,
			session.ID,
			constants.StudyItemVocabulary,
			grammarID,
		)

	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	item := models.StudySessionItem{
		ID:             uuid.New(),
		StudySessionID: session.ID,
		ItemType:       constants.StudyItemGrammar,
		ItemID:         grammarID,
	}

	return s.studySessionRepository.AddItem(
		ctx,
		&item,
	)
}

func (s *studySessionService) UpdateNotes(
	ctx context.Context,
	userID uuid.UUID,
	notes string,
) error {

	session, err :=
		s.getOrCreateTodaySession(
			ctx,
			userID,
		)

	if err != nil {
		return err
	}

	session.Notes = notes

	return s.studySessionRepository.Update(
		ctx,
		session,
	)
}

func (s *studySessionService) UpdateReflection(
	ctx context.Context,
	userID uuid.UUID,
	reflection string,
) error {

	session, err :=
		s.getOrCreateTodaySession(
			ctx,
			userID,
		)

	if err != nil {
		return err
	}

	session.Reflection = reflection

	return s.studySessionRepository.Update(
		ctx,
		session,
	)
}

func (s *studySessionService) GetTodaySession(
	ctx context.Context,
	userID uuid.UUID,
) (*studySessionDTO.TodaySessionResponse, error){
	session, err :=
		s.getOrCreateTodaySession(
			ctx,
			userID,
		)

	if err != nil {
		return nil, err
	}

	items, err :=
	s.studySessionRepository.FindItemsBySessionID(
		ctx,
		session.ID,
	)

	if err != nil {
		return nil, err
	}

	response := &studySessionDTO.TodaySessionResponse{
		Date: session.SessionDate.Format(
			"2006-01-02",
		),
		Notes: session.Notes,
		Reflection: session.Reflection,
		Vocabularies: []studySessionDTO.VocabularyItemResponse{},
		Kanjis: []studySessionDTO.KanjiItemResponse{},
		Grammars: []studySessionDTO.GrammarItemResponse{},
	}

	for _, item := range items {

		switch item.ItemType {

		case constants.StudyItemVocabulary:

			vocabulary, err :=
				s.vocabularyRepository.FindByID(
					ctx,
					item.ItemID,
				)

			if err != nil {
				continue
			}

			response.Vocabularies = append(
				response.Vocabularies,
				studySessionDTO.VocabularyItemResponse{
					ID: vocabulary.ID,
					Word: vocabulary.Word,
					Reading: vocabulary.Reading,
				},
			)

		case constants.StudyItemKanji:

			kanji, err :=
				s.kanjiRepository.FindByID(
					ctx,
					item.ItemID,
				)

			if err != nil {
				continue
			}

			response.Kanjis = append(
				response.Kanjis,
				studySessionDTO.KanjiItemResponse{
					ID: kanji.ID,
					Character: kanji.Character,
					Meaning: kanji.Meaning,
				},
			)

		case constants.StudyItemGrammar:

			grammar, err :=
				s.grammarRepository.FindByID(
					ctx,
					item.ItemID,
				)

			if err != nil {
				continue
			}

			response.Grammars = append(
				response.Grammars,
				studySessionDTO.GrammarItemResponse{
					ID: grammar.ID,
					Pattern: grammar.Pattern,
					Meaning: grammar.Meaning,
				},
			)
		}
	}

	return response, nil
}


