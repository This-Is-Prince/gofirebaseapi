package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}
}

func GetEnv(key string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Env variable %s not found\n", key)
	}
	return env
}
