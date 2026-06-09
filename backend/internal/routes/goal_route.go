package routes

import (
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func GoalRoute(
	router *gin.RouterGroup,
	goalHandler *handlers.GoalHandler,
) {

	goals := router.Group("/goals")

	goals.Use(middlewares.JWTMiddleware())

	goals.POST("",goalHandler.CreateGoal)

	goals.GET("",goalHandler.GetGoalList)

	goals.GET("/:id",goalHandler.GetGoalByID)

	goals.PUT("/:id",goalHandler.UpdateGoal)

	goals.PATCH("/:id/complete",goalHandler.CompleteGoal)

	goals.PATCH("/:id/cancel",goalHandler.CancelGoal)

	goals.DELETE("/:id",goalHandler.DeleteGoal)
}