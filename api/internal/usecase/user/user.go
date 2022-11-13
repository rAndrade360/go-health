package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/rAndrade360/go-health/api/internal/domain"
)

type userUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCase {
	return userUseCase{
		repo: repo,
	}
}

func (u userUseCase) CreateUser(us domain.User) (*domain.User, error) {
	uu, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	us.ID = uu.String()
	us.CreatedAt = time.Now()
	us.UpdatedAt = time.Now()

	err = u.repo.Create(us)
	if err != nil {
		return nil, err
	}

	return &us, nil
}
