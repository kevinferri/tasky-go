package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Println("✅ Env")
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
