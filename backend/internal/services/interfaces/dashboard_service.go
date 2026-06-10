package interfaces

import (
	"context"

	dashboardDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/dashboard"
	"github.com/google/uuid"
)

type DashboardService interface {
	GetDashboard(
		ctx context.Context,
		userID uuid.UUID,
	) (*dashboardDTO.DashboardResponse, error)
}