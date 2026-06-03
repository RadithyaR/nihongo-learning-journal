package main

import (
	"fmt"
	"net/http"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/config"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv();
	database.Connect();
	r := gin.Default();
	port := 8000

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "server is running",
		})
	})

	r.Run(fmt.Sprintf(": %d", port))
}