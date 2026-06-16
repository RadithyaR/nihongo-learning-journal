package routes

import (
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AnalyticsRoute(
	router *gin.RouterGroup,
	analyticsHandler *handlers.AnalyticsHandler,
) {

	analytics := router.Group(
		"/analytics",
	)

	analytics.Use(
		middlewares.JWTMiddleware(),
	)

	analytics.GET(
		"",
		analyticsHandler.GetAnalytics,
	)
}