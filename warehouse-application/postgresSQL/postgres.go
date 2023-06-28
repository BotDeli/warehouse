package postgresSQL

import (
	"database/sql"
	_ "github.com/lib/pq"
	"warehouse-application/errors"
)

func Connect(DbInfo string) *sql.DB {
	db, err := sql.Open("postgres", DbInfo)
	errors.CheckError(err)
	return db
}

func Disconnect(db *sql.DB) {
	err := db.Close()
	errors.CheckWarning(err)
}
