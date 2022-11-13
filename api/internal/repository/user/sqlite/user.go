package repository

import (
	"database/sql"

	"github.com/rAndrade360/go-health/api/internal/domain"
)

type userrepository struct {
	db *sql.DB
}

func NewUserSqliteRepository(db *sql.DB) domain.UserRepository {
	return userrepository{
		db: db,
	}
}

func (u userrepository) Create(user domain.User) error {
	return nil
}
