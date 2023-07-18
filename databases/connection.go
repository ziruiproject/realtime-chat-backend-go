package databases

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"

	_ "github.com/lib/pq"
	_ "github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"
)

func Connection() (*sql.DB, error) {
	helpers.LoadEnv()
	dbData := map[string]string{
		"host": helpers.GetEnv("HOST"),
		"port": helpers.GetEnv("PORT"),
		"user": helpers.GetEnv("DB_USERNAME"),
		"name": helpers.GetEnv("DB_NAME"),
		"pass": helpers.GetEnv("DB_PASSWORD"),
	}

	dbconn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbData["host"], dbData["port"], dbData["user"], dbData["pass"], dbData["name"],
	)

	db, err := sql.Open("postgres", dbconn)
	helpers.ErrorWithLog("Failed to connecting to database", err)
	
	log.Println("Connected to PostreSQL database")

	return db, nil
}
