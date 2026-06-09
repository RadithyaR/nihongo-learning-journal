package goalDTO

type CreateGoalRequest struct {
	Title string `json:"title" validate:"required,max=255"`

	Description *string `json:"description"`

	GoalType *string `json:"goalType"`

	TargetLevel *string `json:"targetLevel"`

	TargetCount *int `json:"targetCount"`

	TargetDate string `json:"targetDate" validate:"required"`
}