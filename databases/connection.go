package databases

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/helper"
)

func Connection() {
	helper.LoadEnv()

	dbData := map[string]string{
		"host": helper.GetEnv("HOST"),
		"port": helper.GetEnv("PORT"),
		"user": helper.GetEnv("DB_USERNAME"),
		"name": helper.GetEnv("DB_NAME"),
		"pass": helper.GetEnv("DB_PASSWORD"),
	}

	dbconn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbData["host"], dbData["port"], dbData["user"], dbData["pass"], dbData["name"],
	)

	db, err := sql.Open("postgres", dbconn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to PostreSQL database")
}
