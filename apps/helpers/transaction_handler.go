package helpers

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		error := tx.Rollback()
		ErrorWithLog("Failed to rollback the transaction", error)
		panic(err)
	} else {
		tx.Commit()
	}
}
