package databases

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"
	"github.com/ziruiproject/realtime-chat-backend-go/apps/configuration"
	_ "github.com/lib/pq"
	_ "github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"
)

func Connection() (*sql.DB, error) {

	dbData := map[string]string{
		"host": configuration.EnvConfigs.DBHost,
		"port": configuration.EnvConfigs.DBPort,
		"user": configuration.EnvConfigs.DBUsername,
		"name": configuration.EnvConfigs.DBName,
		"pass": configuration.EnvConfigs.DBPassword,
	}

	dbconn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbData["host"], dbData["port"], dbData["user"], dbData["pass"], dbData["name"],
	)

	db, err := sql.Open("postgres", dbconn)
	helpers.ErrorWithLog("Failed to connecting to database", err)
	
	log.Println("Connected to PostreSQL database")
	log.Println(configuration.EnvConfigs.DBHost)

	return db, nil
}
