package interfaces

import (
	"context"

	tagDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/tag"
	"github.com/google/uuid"
)

type TaggableService interface {
	AttachTag(
		ctx context.Context,
		userID uuid.UUID,
		dto tagDTO.AttachTagRequest,
	) error

	RemoveTag(
		ctx context.Context,
		userID uuid.UUID,
		dto tagDTO.RemoveTagRequest,
	) error

	GetTagsByItem(
		ctx context.Context,
		userID uuid.UUID,
		itemType string,
		itemID uuid.UUID,
	) ([]tagDTO.TagResponse, error)
}