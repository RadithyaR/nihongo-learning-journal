package goalDTO

import "github.com/google/uuid"

type GoalResponse struct {
	ID uuid.UUID `json:"id"`

	Title string `json:"title"`

	Description *string `json:"description"`

	GoalType *string `json:"goalType"`

	TargetLevel *string `json:"targetLevel"`

	TargetCount *int `json:"targetCount"`

	CurrentCount int `json:"currentCount"`

	ProgressPercentage float64 `json:"progressPercentage"`

	TargetDate string `json:"targetDate"`

	Status string `json:"status"`
}