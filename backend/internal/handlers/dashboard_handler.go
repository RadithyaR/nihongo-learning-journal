package handlers

import (
	"net/http"

	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	dashboardService serviceInterfaces.DashboardService
}

func NewDashboardHandler(
	dashboardService serviceInterfaces.DashboardService,
) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
	}
}

func (h *DashboardHandler) GetDashboard(
	c *gin.Context,
) {

	userID, ok := GetUserID(c)

	if !ok {
		responses.Error(
			c,
			http.StatusUnauthorized,
			"user not found",
		)
		return
	}

	dashboard, err := h.dashboardService.GetDashboard(
		c.Request.Context(),
		userID,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusInternalServerError,
			"failed to retrieve dashboard",
		)

		return
	}

	responses.Success(
		c,
		http.StatusOK,
		"dashboard retrieved successfully",
		dashboard,
	)
}