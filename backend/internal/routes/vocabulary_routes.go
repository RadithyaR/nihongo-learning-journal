package routes

import (
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func VocabularyRoute(
	router *gin.RouterGroup,
	vocabularyHandler *handlers.VocabularyHandler,
) {

	vocabularies := router.Group("/vocabularies")

	vocabularies.Use(middlewares.JWTMiddleware())

	vocabularies.POST("", vocabularyHandler.CreateVocabulary)

	vocabularies.GET("", vocabularyHandler.GetVocabularyList)

	vocabularies.GET("/:id", vocabularyHandler.GetVocabularyByID)

	vocabularies.PUT("/:id", vocabularyHandler.UpdateVocabulary)

	vocabularies.DELETE("/:id", vocabularyHandler.DeleteVocabulary)
}