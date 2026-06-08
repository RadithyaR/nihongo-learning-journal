package routes

import (
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func KanjiRoute(
	router *gin.RouterGroup,
	kanjiHandler *handlers.KanjiHandler,
) {

	kanji := router.Group("/kanjis",middlewares.JWTMiddleware())

	kanji.POST("", kanjiHandler.CreateKanji)

	kanji.GET("", kanjiHandler.GetKanjiList)

	kanji.GET("/:id", kanjiHandler.GetKanjiByID)

	kanji.PUT("/:id", kanjiHandler.UpdateKanji)

	kanji.DELETE("/:id", kanjiHandler.DeleteKanji)

	kanji.PATCH("/:id/favourite", kanjiHandler.ToggleFavourite)
}