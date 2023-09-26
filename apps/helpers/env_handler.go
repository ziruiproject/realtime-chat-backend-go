package helpers

import (
	"github.com/gofor-little/env"
	configs "github.com/ziruiproject/realtime-chat-backend-go/apps/configuration"
	"log"
	//	"github.com/joho/godotenv"
	//	_ "github.com/joho/godotenv/autoload"
	//	"path/filepath"
)

func LoadEnv() {
	err := env.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	 configs.EnvConfigs = &configs.Config{
		DBUsername: GetEnv("DB_USERNAME"),
		DBPassword: GetEnv("DB_PASSWORD"),
		DBName: GetEnv("DB_NAME"),
		DBPort: GetEnv("PORT"),
		DBHost: GetEnv("HOST"),
		SecretToken: GetEnv("JWT_SECRET"),
	}
}

func GetEnv(name string) string {
	value, err := env.MustGet(name)
	if err != nil {
		log.Fatal(err)
	}
	return value
}