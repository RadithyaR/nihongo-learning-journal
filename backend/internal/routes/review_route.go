package routes

import (
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func ReviewRoute(
	router *gin.RouterGroup,
	reviewHandler *handlers.ReviewHandler,
) {

	reviews := router.Group("/reviews")

	reviews.Use(middlewares.JWTMiddleware())

	reviews.GET("/next", reviewHandler.GetNextReview)
	reviews.POST("", reviewHandler.SubmitReview)

	reviews.GET("/kanji/next", reviewHandler.GetNextKanjiReview)
	reviews.POST("/kanji", reviewHandler.SubmitKanjiReview)

	reviews.GET("/grammar/next", reviewHandler.GetNextGrammarReview)
	reviews.POST("/grammar", reviewHandler.SubmitGrammarReview)
}