package sqlite

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() (*sql.DB, error) {
	dbpath := os.Getenv("SQLITE_DB_PATH")
	if dbpath == "" {
		return nil, errors.New("db path not set")
	}

	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
