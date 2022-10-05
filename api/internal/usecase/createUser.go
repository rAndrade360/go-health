package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/rAndrade360/go-health/api/internal/model"
	repo "github.com/rAndrade360/go-health/api/internal/repository/user"
)

func CreateUser(u model.User) (*model.User, error) {
	uu, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	u.ID = uu.String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	err = repo.Create(u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
