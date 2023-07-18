package databases

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"

	_ "github.com/lib/pq"
	_ "github.com/ziruiproject/realtime-chat-backend-go/apps/helpers"
)

type Database struct {
	Db *sql.DB
	Tx *sql.Tx
}

func ConstrDatabase() *Database {
	var db *sql.DB = connection()
	var tx *sql.Tx = txConnection(db)

	return &Database{
		Db: db,
		Tx: tx,
	}
}

func connection() *sql.DB {
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

	defer func(db *sql.DB) {
		err := db.Close()
		helpers.ErrorWithLog("Failed to close db connection", err)
	}(db)

	log.Println("Connected to PostreSQL database")

	return db
}

func txConnection(db *sql.DB) *sql.Tx {

	tx, err := db.Begin()
	helpers.ErrorWithLog("Failed to start transaction:", err)

	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		helpers.ErrorWithLog("Failed to rollback transaction:", err)
	}(tx)

	err = tx.Commit()
	helpers.ErrorWithLog("Failed to commit transaction:", err)

	return tx
}
