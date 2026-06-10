package services

import (
	"context"
	"errors"

	tagDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/tag"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/google/uuid"
)

type tagService struct {
	tagRepo repositoryInterfaces.TagRepository
}

func NewTagService(
	tagRepo repositoryInterfaces.TagRepository,
) serviceInterfaces.TagService {
	return &tagService{
		tagRepo: tagRepo,
	}
}

func (s *tagService) CreateTag(
	ctx context.Context,
	userID uuid.UUID,
	dto tagDTO.CreateTagRequest,
) (*tagDTO.TagResponse, error) {

	exists, err := s.tagRepo.ExistsByName(
		ctx,
		userID,
		dto.Name,
	)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New(
			"tag already exists",
		)
	}

	tag := models.Tag{
		UserID: userID,
		Name:   dto.Name,
		Color:  dto.Color,
	}

	if err := s.tagRepo.Create(
		ctx,
		&tag,
	); err != nil {
		return nil, err
	}

	return &tagDTO.TagResponse{
		ID:        tag.ID.String(),
		Name:      tag.Name,
		Color:     tag.Color,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
	}, nil
}

func (s *tagService) GetTags(
	ctx context.Context,
	userID uuid.UUID,
) ([]tagDTO.TagResponse, error) {

	tags, err := s.tagRepo.FindByUserID(
		ctx,
		userID,
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

func (s *tagService) UpdateTag(
	ctx context.Context,
	userID uuid.UUID,
	tagID uuid.UUID,
	dto tagDTO.UpdateTagRequest,
) (*tagDTO.TagResponse, error) {

	tag, err := s.tagRepo.FindByID(
		ctx,
		tagID,
	)

	if err != nil {
		return nil, err
	}

	if tag.UserID != userID {
		return nil, errors.New(
			"tag not found",
		)
	}

	tag.Name = dto.Name
	tag.Color = dto.Color

	if err := s.tagRepo.Update(
		ctx,
		tag,
	); err != nil {
		return nil, err
	}

	return &tagDTO.TagResponse{
		ID:        tag.ID.String(),
		Name:      tag.Name,
		Color:     tag.Color,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
	}, nil
}

func (s *tagService) DeleteTag(
	ctx context.Context,
	userID uuid.UUID,
	tagID uuid.UUID,
) error {

	tag, err := s.tagRepo.FindByID(
		ctx,
		tagID,
	)

	if err != nil {
		return err
	}

	if tag.UserID != userID {
		return errors.New(
			"tag not found",
		)
	}

	return s.tagRepo.Delete(
		ctx,
		tag,
	)
}