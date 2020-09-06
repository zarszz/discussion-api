package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase - connect and create migration
func ConnectDatabase() {
	dsn := GoDotEnvVariable("CONN_STRING")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Comment{})
	database.AutoMigrate(&Discussion{})
	database.AutoMigrate(&Category{})
	DB = database
}

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
