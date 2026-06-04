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

	userRepository := repositories.NewUserRepository(
		database.DB,
	)
	userSessionRepository :=
	repositories.NewUserSessionRepository(
		database.DB,
	)

	authService := services.NewAuthService(
		userRepository,
		userSessionRepository,
	)

	authHandler := handlers.NewAuthHandler(
		authService,
		
	)

	//route
	api := r.Group("/api/v1")	

	routes.AuthRoute(api, authHandler,)

	r.Run(fmt.Sprintf(":%d", port))
}