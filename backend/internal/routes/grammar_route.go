package routes

import (
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func GrammarRoute(
	router *gin.RouterGroup,
	grammarHandler *handlers.GrammarHandler,
) {

	grammars := router.Group("/grammars")

	grammars.Use(middlewares.JWTMiddleware())

	grammars.POST("", grammarHandler.CreateGrammar)

	grammars.GET("", grammarHandler.GetGrammarList)

	grammars.GET("/:id", grammarHandler.GetGrammarByID)

	grammars.PUT("/:id", grammarHandler.UpdateGrammar)

	grammars.DELETE("/:id", grammarHandler.DeleteGrammar,)

	grammars.PATCH("/:id/favourite", grammarHandler.ToggleFavourite)
}