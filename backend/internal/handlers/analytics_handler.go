package handlers

import (
	"net/http"

	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	analyticsService serviceInterfaces.AnalyticsService
}

func NewAnalyticsHandler(
	analyticsService serviceInterfaces.AnalyticsService,
) *AnalyticsHandler {
	return &AnalyticsHandler{
		analyticsService: analyticsService,
	}
}

func (h *AnalyticsHandler) GetAnalytics(
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

	data, err := h.analyticsService.GetAnalytics(
		c.Request.Context(),
		userID,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)

		return
	}

	responses.Success(
		c,
		http.StatusOK,
		"analytics retrieved successfully",
		data,
	)
}