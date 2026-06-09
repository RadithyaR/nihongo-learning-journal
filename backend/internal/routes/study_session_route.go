package routes

import (
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func StudySessionRoute(
	router *gin.RouterGroup,
	studySessionHandler *handlers.StudySessionHandler,
) {

	studySessions := router.Group("/study-sessions")

	studySessions.Use(middlewares.JWTMiddleware())

	studySessions.GET("/today", studySessionHandler.GetTodaySession)

	studySessions.PATCH("/notes", studySessionHandler.UpdateNotes)

	studySessions.PATCH("/reflection", studySessionHandler.UpdateReflection)
}