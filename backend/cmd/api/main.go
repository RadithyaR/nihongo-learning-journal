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
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/email"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/validator"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv();
	database.Connect();
	validator.Init()
	r := gin.Default();
	
	// Setup CORS
	frontendURL := config.GetEnv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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

	srsRepository :=
	repositories.NewSRSRepository(database.DB)

	goalRepository :=
	repositories.NewGoalRepository(database.DB)

	dashboardRepository := 
	repositories.NewDashboardRepository(database.DB)

	tagRepository := 
	repositories.NewTagRepository(database.DB)

	taggableRepository :=
	repositories.NewTaggableRepository(database.DB)

	analyticsRepository :=
	repositories.NewAnalyticsRepository(database.DB)

	//service
	emailService := email.NewEmailService()

	authService := services.NewAuthService(
		userRepository,
		userSessionRepository,
		emailVerificationRepository,
		passwordResetRepository,
		emailService,
	)

	profileService := services.NewProfileService(userRepository)

	vocabularyService := services.NewVocabularyService(vocabularyRepository, srsRepository)

	studySessionService := services.NewStudySessionService(
	studySessionRepository,
	vocabularyRepository,
	kanjiRepository,
	grammarRepository,
)

	srsService := services.NewSRSService(srsRepository)

	reviewService := services.NewReviewService(
		reviewRepository,
		vocabularyRepository,
		kanjiRepository,
		grammarRepository,
		studySessionService,
		srsService,
	)

	kanjiService := services.NewKanjiService(kanjiRepository, srsRepository)

	grammarService := services.NewGrammarService(grammarRepository, srsRepository)

	goalService := services.NewGoalService(goalRepository)

	dashboardService := services.NewDashboardService(
		dashboardRepository,
		srsService,
	)

	tagService := services.NewTagService(tagRepository)

	taggableService := services.NewTaggableService(
		taggableRepository,
		tagRepository,
		vocabularyRepository,
		kanjiRepository,
		grammarRepository,
	)

	analyticsService := services.NewAnalyticsService(analyticsRepository)

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

	analyticsHandler := handlers.NewAnalyticsHandler(analyticsService)

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

	routes.AnalyticsRoute(api, analyticsHandler)

	r.Run(fmt.Sprintf(":%d", port))
}