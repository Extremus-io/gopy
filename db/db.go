package db

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"github.com/Extremus-io/gopy/log"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("sqlite3", "file:gopy.db?cache=shared&mode=memory&_loc=auto")
	if err != nil {
		panic(err)
	}
	log.Verbose("Database connection successful")
}