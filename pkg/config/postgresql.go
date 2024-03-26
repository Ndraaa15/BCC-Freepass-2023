package config

import (
	"fmt"
	"os"
)

func LoadConfigPostgresql() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_max_conns=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
		os.Getenv("DB_POOL_MAX_CONNS"),
	)
}
