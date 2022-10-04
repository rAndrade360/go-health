package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/rAndrade360/go-health/api/internal/model"
	repo "github.com/rAndrade360/go-health/api/internal/repository/user"
)

func CreateUser(u *model.User) error {
	uu, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	u.ID = uu.String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return repo.Create(*u)
}
