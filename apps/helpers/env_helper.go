package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
}

func GetEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatal("Env not found")
	}
	return value
}
