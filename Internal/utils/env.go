package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		// Running in container or env-only mode, no need to spam logs
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Println("Warning : unable to load .env file: &v", err)
	} else {
		log.Println("Environment variables loaded successfully.")
	}

}
