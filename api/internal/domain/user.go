package domain

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	Weight    float64   `json:"weight"`
	BirthDate string    `json:"birthdate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

const (
	PENDING  = "PENDING"
	ACTIVE   = "ACTIVE"
	INACTIVE = "INACTIVE"
)

type UserRepository interface {
	Create(user User) error
}

type UserUseCase interface {
	CreateUser(us User) (*User, error)
}
