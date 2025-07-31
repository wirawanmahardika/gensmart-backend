package config

import (
	"os"

	"github.com/joho/godotenv"
)

var AppConfig *Config

type Config struct {
	DBName string
	DBUser string
	DBPass string

	Port string
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic("failed load .env file")
	}

	AppConfig = &Config{
		Port: os.Getenv("PORT"),

		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
	}
}
