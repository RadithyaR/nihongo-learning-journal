package routes

import (
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func ProfileRoute(
	router *gin.RouterGroup,
	profileHandler *handlers.ProfileHandler,
) {

	profile := router.Group("/profile")

	profile.Use(
		middlewares.JWTMiddleware(),
	)

	profile.GET(
		"",
		profileHandler.GetProfile,
	)

	profile.PUT(
		"",
		profileHandler.UpdateProfile,
	)
}