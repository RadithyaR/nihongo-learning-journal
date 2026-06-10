package services

import (
	"context"
	"errors"
	"time"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/constants"
	goalDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/goal"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterface "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type goalService struct {
	goalRepository repositoryInterfaces.GoalRepository
}

func NewGoalService(
	goalRepository repositoryInterfaces.GoalRepository,
) serviceInterface.GoalService {

	return &goalService{
		goalRepository: goalRepository,
	}
}

func (s *goalService) mapGoalResponse(
	goal *models.Goal,
	currentCount int,
	progressPercentage float64,
) *goalDTO.GoalResponse {

	return &goalDTO.GoalResponse{
		ID:                 goal.ID,
		Title:              goal.Title,
		Description:        goal.Description,
		GoalType:           goal.GoalType,
		TargetLevel:        goal.TargetLevel,
		TargetCount:        goal.TargetCount,
		CurrentCount:       currentCount,
		ProgressPercentage: progressPercentage,
		TargetDate:         goal.TargetDate.Format("2006-01-02"),
		Status:             goal.Status,
	}
}

func (s *goalService) calculateProgress(
	ctx context.Context,
	userID uuid.UUID,
	goal *models.Goal,
) (int, float64, error) {

	if goal.GoalType == nil ||
		goal.TargetCount == nil {

		return 0, 0, nil
	}

	count, err :=
		s.goalRepository.CountProgress(
			ctx,
			userID,
			*goal.GoalType,
		)

	if err != nil {
		return 0, 0, err
	}

	progress := 0.0

	if *goal.TargetCount > 0 {

		progress =
			(float64(count) /
				float64(*goal.TargetCount)) * 100

		if progress > 100 {
			progress = 100
		}
	}

	return int(count), progress, nil
}

func (s *goalService) CreateGoal(
	ctx context.Context,
	userID uuid.UUID,
	dto goalDTO.CreateGoalRequest,
) (*goalDTO.GoalResponse, error) {

	targetDate, err := time.Parse(
		"2006-01-02",
		dto.TargetDate,
	)

	if err != nil {
		return nil, err
	}

	goal := models.Goal{
		UserID:      userID,
		Title:       dto.Title,
		Description: dto.Description,
		GoalType:    dto.GoalType,
		TargetLevel: dto.TargetLevel,
		TargetCount: dto.TargetCount,
		TargetDate:  targetDate,
		Status:      constants.GoalStatusInProgress,
	}

	if err := s.goalRepository.Create(
		ctx,
		&goal,
	); err != nil {
		return nil, err
	}

	return s.mapGoalResponse(
		&goal,
		0,
		0,
	), nil
}

func (s *goalService) GetGoalByID(
	ctx context.Context,
	userID uuid.UUID,
	goalID uuid.UUID,
) (*goalDTO.GoalResponse, error) {

	goal, err :=
		s.goalRepository.FindByID(
			ctx,
			goalID,
		)

	if err != nil {

		if errors.Is(
			err,
			gorm.ErrRecordNotFound,
		) {
			return nil, customErrors.ErrGoalNotFound
		}

		return nil, err
	}

	if goal.UserID != userID {
		return nil, customErrors.ErrGoalNotFound
	}

	currentCount,
		progress,
		err :=
		s.calculateProgress(
			ctx,
			userID,
			goal,
		)

	if err != nil {
		return nil, err
	}

	return s.mapGoalResponse(
		goal,
		currentCount,
		progress,
	), nil
}

func (s *goalService) GetGoalList(
	ctx context.Context,
	userID uuid.UUID,
) ([]goalDTO.GoalResponse, error) {

	goals, err :=
		s.goalRepository.FindAllByUserID(
			ctx,
			userID,
		)

	if err != nil {
		return nil, err
	}

	responses :=
		make(
			[]goalDTO.GoalResponse,
			0,
			len(goals),
		)

	for _, goal := range goals {

		currentCount,
			progress,
			err :=
			s.calculateProgress(
				ctx,
				userID,
				&goal,
			)

		if err != nil {
			return nil, err
		}

		responses = append(
			responses,
			* s.mapGoalResponse(
				&goal,
				currentCount,
				progress,
			),
		)
	}

	return responses, nil
}

func (s *goalService) UpdateGoal(
	ctx context.Context,
	userID uuid.UUID,
	goalID uuid.UUID,
	dto goalDTO.UpdateGoalRequest,
) (*goalDTO.GoalResponse, error) {

	goal, err :=
		s.goalRepository.FindByID(
			ctx,
			goalID,
		)

	if err != nil {

		if errors.Is(
			err,
			gorm.ErrRecordNotFound,
		) {
			return nil,
				customErrors.ErrGoalNotFound
		}

		return nil, err
	}

	if goal.UserID != userID {
		return nil,
			customErrors.ErrGoalNotFound
	}

	targetDate, err := time.Parse(
		"2006-01-02",
		dto.TargetDate,
	)

	if err != nil {
		return nil, err
	}

	goal.Title = dto.Title
	goal.Description = dto.Description
	goal.GoalType = dto.GoalType
	goal.TargetLevel = dto.TargetLevel
	goal.TargetCount = dto.TargetCount
	goal.TargetDate = targetDate

	if err := s.goalRepository.Update(
		ctx,
		goal,
	); err != nil {
		return nil, err
	}

	currentCount,
		progress,
		err :=
		s.calculateProgress(
			ctx,
			userID,
			goal,
		)

	if err != nil {
		return nil, err
	}

	return s.mapGoalResponse(
		goal,
		currentCount,
		progress,
	), nil
}

func (s *goalService) CompleteGoal(
	ctx context.Context,
	userID uuid.UUID,
	goalID uuid.UUID,
) error {

	goal, err :=
		s.goalRepository.FindByID(
			ctx,
			goalID,
		)

	if err != nil {

		if errors.Is(
			err,
			gorm.ErrRecordNotFound,
		) {
			return customErrors.ErrGoalNotFound
		}

		return err
	}

	if goal.UserID != userID {
		return customErrors.ErrGoalNotFound
	}

	goal.Status =
		constants.GoalStatusCompleted

	return s.goalRepository.Update(
		ctx,
		goal,
	)
}

func (s *goalService) CancelGoal(
	ctx context.Context,
	userID uuid.UUID,
	goalID uuid.UUID,
) error {

	goal, err :=
		s.goalRepository.FindByID(
			ctx,
			goalID,
		)

	if err != nil {

		if errors.Is(
			err,
			gorm.ErrRecordNotFound,
		) {
			return customErrors.ErrGoalNotFound
		}

		return err
	}

	if goal.UserID != userID {
		return customErrors.ErrGoalNotFound
	}

	goal.Status =
		constants.GoalStatusCancelled

	return s.goalRepository.Update(
		ctx,
		goal,
	)
}

func (s *goalService) DeleteGoal(
	ctx context.Context,
	userID uuid.UUID,
	goalID uuid.UUID,
) error {

	goal, err :=
		s.goalRepository.FindByID(
			ctx,
			goalID,
		)

	if err != nil {

		if errors.Is(
			err,
			gorm.ErrRecordNotFound,
		) {
			return customErrors.ErrGoalNotFound
		}

		return err
	}

	if goal.UserID != userID {
		return customErrors.ErrGoalNotFound
	}

	return s.goalRepository.Delete(
		ctx,
		goalID,
	)
}