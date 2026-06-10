package routes

import (
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func DashboardRoute(
	router *gin.RouterGroup,
	dashboardHandler *handlers.DashboardHandler,
) {

	dashboard := router.Group("/dashboard")

	dashboard.Use(middlewares.JWTMiddleware())

	dashboard.GET("", dashboardHandler.GetDashboard)
}