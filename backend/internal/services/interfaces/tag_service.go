package interfaces

import (
	"context"

	tagDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/tag"
	"github.com/google/uuid"
)

type TagService interface {
	CreateTag(
		ctx context.Context,
		userID uuid.UUID,
		dto tagDTO.CreateTagRequest,
	) (*tagDTO.TagResponse, error)

	GetTags(
		ctx context.Context,
		userID uuid.UUID,
	) ([]tagDTO.TagResponse, error)

	UpdateTag(
		ctx context.Context,
		userID uuid.UUID,
		tagID uuid.UUID,
		dto tagDTO.UpdateTagRequest,
	) (*tagDTO.TagResponse, error)

	DeleteTag(
		ctx context.Context,
		userID uuid.UUID,
		tagID uuid.UUID,
	) error
}