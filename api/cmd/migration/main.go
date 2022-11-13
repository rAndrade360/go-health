package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rAndrade360/go-health/api/adapter/database/sqlite"
	"github.com/rAndrade360/go-health/api/internal/database/migration"
)

func init() {
	godotenv.Load("../../.env")
}

func main() {
	db, err := sqlite.GetDB()
	if err != nil {
		log.Fatalf("Error to get DB: %s", err.Error())
	}

	migration.CreateUserTable(db)
}
