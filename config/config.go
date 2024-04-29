package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	mysql "go-wishlist-api/drivers/mysql"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitConfigMySQL() mysql.Config {
	return mysql.Config{
		DBName:     os.Getenv("DB_NAME"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPass:     os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}
