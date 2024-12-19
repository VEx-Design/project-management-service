package pkg

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}
}
