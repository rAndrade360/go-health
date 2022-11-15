package repository

import (
	"database/sql"
	"time"

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
	stmt, err := u.db.Prepare("INSERT INTO user VALUES (?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.ID, user.Name, user.Username, user.Email, user.Status, user.Weight, user.BirthDate, user.CreatedAt.Format(time.RFC3339), user.UpdatedAt.Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}
