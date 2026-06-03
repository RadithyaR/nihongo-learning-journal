package database

import (
	"log"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	dsn := config.GetEnv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	DB = db

	log.Println("Database connected successfuly")
}