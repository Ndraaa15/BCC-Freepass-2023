package config

import (
	"fmt"
	"os"
)

func LoadConfigPostgresql() string {
	return fmt.Sprintf("user=%s dbname=%s sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), "verify-full")
}
