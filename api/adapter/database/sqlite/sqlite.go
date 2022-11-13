package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "temp.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
