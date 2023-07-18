package helpers

import (
	"database/sql"
	"log"
)

func ErrorWithLog(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func ErrorDefer(rows *sql.Rows) {
	err := rows.Close()
	ErrorWithLog("Failed closing connection", err)
}
