package interfaces

import (
	"context"

	goalDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/goal"
	"github.com/google/uuid"
)

type GoalService interface {
	CreateGoal(
		ctx context.Context,
		userID uuid.UUID,
		dto goalDTO.CreateGoalRequest,
	) (*goalDTO.GoalResponse, error)

	GetGoalList(
		ctx context.Context,
		userID uuid.UUID,
	) ([]goalDTO.GoalResponse, error)

	GetGoalByID(
		ctx context.Context,
		userID uuid.UUID,
		goalID uuid.UUID,
	) (*goalDTO.GoalResponse, error)

	UpdateGoal(
		ctx context.Context,
		userID uuid.UUID,
		goalID uuid.UUID,
		dto goalDTO.UpdateGoalRequest,
	) (*goalDTO.GoalResponse, error)

	CompleteGoal(
		ctx context.Context,
		userID uuid.UUID,
		goalID uuid.UUID,
	) error

	CancelGoal(
		ctx context.Context,
		userID uuid.UUID,
		goalID uuid.UUID,
	) error

	DeleteGoal(
		ctx context.Context,
		userID uuid.UUID,
		goalID uuid.UUID,
	) error
}