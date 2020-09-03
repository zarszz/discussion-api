package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/zarszz/discussion-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase() {
	dsn := getDotEnv("CONNECTION_STRING")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Comment{})
	database.AutoMigrate(&models.Discussions{})
	DB = database
}

func getDotEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return ""
	}
	return os.Getenv(key)
}
