package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	return nil
}
