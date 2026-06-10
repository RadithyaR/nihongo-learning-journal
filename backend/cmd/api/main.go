package main

import (
	"fmt"
	"net/http"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/config"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/database"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/routes"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/services"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/validator"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv();
	database.Connect();
	validator.Init()
	r := gin.Default();
	port := 8000

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "server is running",
		})
	})

	//repository
	userRepository := repositories.NewUserRepository(database.DB)
	userSessionRepository := repositories.NewUserSessionRepository(database.DB)

	emailVerificationRepository := 
	repositories.NewEmailVerificationRepository(database.DB)

	passwordResetRepository := 
	repositories.NewPasswordResetRepository(database.DB)

	vocabularyRepository := 
	repositories.NewVocabularyRepository(database.DB)

	reviewRepository := 
	repositories.NewReviewRepository(database.DB)

	kanjiRepository :=
	repositories.NewKanjiRepository(database.DB)

	grammarRepository :=
	repositories.NewGrammarRepository(database.DB)

	studySessionRepository :=
	repositories.NewStudySessionRepository(database.DB)

	goalRepository :=
	repositories.NewGoalRepository(database.DB)

	dashboardRepository := 
	repositories.NewDashboardRepository(database.DB)

	tagRepository := 
	repositories.NewTagRepository(database.DB)

	taggableRepository :=
	repositories.NewTaggableRepository(database.DB)

	//service
	authService := services.NewAuthService(
		userRepository,
		userSessionRepository,
		emailVerificationRepository,
		passwordResetRepository,
	)

	profileService := services.NewProfileService(userRepository)

	vocabularyService := services.NewVocabularyService(vocabularyRepository)

	studySessionService := services.NewStudySessionService(
	studySessionRepository,
	vocabularyRepository,
	kanjiRepository,
	grammarRepository,
)

	reviewService := services.NewReviewService(
		reviewRepository,
		vocabularyRepository,
		kanjiRepository,
		grammarRepository,
		studySessionService,
	)

	kanjiService := services.NewKanjiService(kanjiRepository)

	grammarService := services.NewGrammarService(grammarRepository)

	goalService := services.NewGoalService(goalRepository)

	dashboardService := services.NewDashboardService(dashboardRepository)

	tagService := services.NewTagService(tagRepository)

	taggableService := services.NewTaggableService(
		taggableRepository,
		tagRepository,
		vocabularyRepository,
		kanjiRepository,
		grammarRepository,
	)

	//handler
	authHandler := handlers.NewAuthHandler(authService)

	profileHandler := handlers.NewProfileHandler(profileService)

	vocabularyHandler :=handlers.NewVocabularyHandler(vocabularyService)

	reviewHandler := handlers.NewReviewHandler(reviewService)

	kanjiHandler := handlers.NewKanjiHandler(kanjiService)

	grammarHandler := handlers.NewGrammarHandler(grammarService)

	studySessionHandler := handlers.NewStudySessionHandler(studySessionService)

	goalHandler := handlers.NewGoalHandler(goalService)

	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	tagHandler := handlers.NewTagHandler(tagService)

	taggableHandler := handlers.NewTaggableHandler(taggableService)

	//route
	api := r.Group("/api/v1")	

	routes.AuthRoute(api, authHandler)

	routes.ProfileRoute(api, profileHandler)

	routes.VocabularyRoute(api, vocabularyHandler)
	
	routes.ReviewRoute(api, reviewHandler)

	routes.KanjiRoute(api, kanjiHandler)

	routes.GrammarRoute(api, grammarHandler)

	routes.StudySessionRoute(api, studySessionHandler)

	routes.GoalRoute(api, goalHandler)

	routes.DashboardRoute(api, dashboardHandler)

	routes.TagRoute(api, tagHandler, taggableHandler)

	r.Run(fmt.Sprintf(":%d", port))
}