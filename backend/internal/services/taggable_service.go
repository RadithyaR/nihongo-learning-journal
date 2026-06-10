package services

import (
	"context"
	"errors"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/constants"
	tagDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/tag"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/google/uuid"
)

type taggableService struct {
	taggableRepo   repositoryInterfaces.TaggableRepository
	tagRepo        repositoryInterfaces.TagRepository
	vocabularyRepo repositoryInterfaces.VocabularyRepository
	kanjiRepo      repositoryInterfaces.KanjiRepository
	grammarRepo    repositoryInterfaces.GrammarRepository
}

func NewTaggableService(
	taggableRepo repositoryInterfaces.TaggableRepository,
	tagRepo repositoryInterfaces.TagRepository,
	vocabularyRepo repositoryInterfaces.VocabularyRepository,
	kanjiRepo repositoryInterfaces.KanjiRepository,
	grammarRepo repositoryInterfaces.GrammarRepository,
) serviceInterfaces.TaggableService {

	return &taggableService{
		taggableRepo:   taggableRepo,
		tagRepo:        tagRepo,
		vocabularyRepo: vocabularyRepo,
		kanjiRepo:      kanjiRepo,
		grammarRepo:    grammarRepo,
	}
}

func (s *taggableService) validateOwnership(
	ctx context.Context,
	userID uuid.UUID,
	itemType string,
	itemID uuid.UUID,
) error {

	switch itemType {

	case constants.ItemTypeVocabulary:

		vocabulary, err := s.vocabularyRepo.FindByID(
			ctx,
			itemID,
		)

		if err != nil {
			return err
		}

		if vocabulary.UserID != userID {
			return errors.New("item not found")
		}

	case constants.ItemTypeKanji:

		kanji, err := s.kanjiRepo.FindByID(
			ctx,
			itemID,
		)

		if err != nil {
			return err
		}

		if kanji.UserID != userID {
			return errors.New("item not found")
		}

	case constants.ItemTypeGrammar:

		grammar, err := s.grammarRepo.FindByID(
			ctx,
			itemID,
		)

		if err != nil {
			return err
		}

		if grammar.UserID != userID {
			return errors.New("item not found")
		}

	default:
		return errors.New("invalid item type")
	}

	return nil
}

func (s *taggableService) AttachTag(
	ctx context.Context,
	userID uuid.UUID,
	dto tagDTO.AttachTagRequest,
) error {

	tagID, err := uuid.Parse(dto.TagID)
	if err != nil {
		return err
	}

	itemID, err := uuid.Parse(dto.ItemID)
	if err != nil {
		return err
	}

	tag, err := s.tagRepo.FindByID(
		ctx,
		tagID,
	)

	if err != nil {
		return err
	}

	if tag.UserID != userID {
		return errors.New("tag not found")
	}

	if err := s.validateOwnership(
		ctx,
		userID,
		dto.ItemType,
		itemID,
	); err != nil {
		return err
	}

	exists, err := s.taggableRepo.Exists(
		ctx,
		tagID,
		dto.ItemType,
		itemID,
	)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("tag already attached")
	}

	taggable := models.Taggable{
		TagID:    tagID,
		ItemType: dto.ItemType,
		ItemID:   itemID,
	}

	return s.taggableRepo.Create(
		ctx,
		&taggable,
	)
}

func (s *taggableService) RemoveTag(
	ctx context.Context,
	userID uuid.UUID,
	dto tagDTO.RemoveTagRequest,
) error {

	tagID, err := uuid.Parse(dto.TagID)
	if err != nil {
		return err
	}

	itemID, err := uuid.Parse(dto.ItemID)
	if err != nil {
		return err
	}

	tag, err := s.tagRepo.FindByID(
		ctx,
		tagID,
	)

	if err != nil {
		return err
	}

	if tag.UserID != userID {
		return errors.New("tag not found")
	}

	if err := s.validateOwnership(
		ctx,
		userID,
		dto.ItemType,
		itemID,
	); err != nil {
		return err
	}

	return s.taggableRepo.Delete(
		ctx,
		tagID,
		dto.ItemType,
		itemID,
	)
}

func (s *taggableService) GetTagsByItem(
	ctx context.Context,
	userID uuid.UUID,
	itemType string,
	itemID uuid.UUID,
) ([]tagDTO.TagResponse, error) {

	if err := s.validateOwnership(
		ctx,
		userID,
		itemType,
		itemID,
	); err != nil {
		return nil, err
	}

	tags, err := s.taggableRepo.FindTagsByItem(
		ctx,
		itemType,
		itemID,
	)

	if err != nil {
		return nil, err
	}

	responses := make(
		[]tagDTO.TagResponse,
		0,
		len(tags),
	)

	for _, tag := range tags {

		responses = append(
			responses,
			tagDTO.TagResponse{
				ID:        tag.ID.String(),
				Name:      tag.Name,
				Color:     tag.Color,
				CreatedAt: tag.CreatedAt,
				UpdatedAt: tag.UpdatedAt,
			},
		)
	}

	return responses, nil
}