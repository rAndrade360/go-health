package domain

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Weight    float64   `json:"weight"`
	BirthDate string    `json:"birthdate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserRepository interface {
	Create(user User) error
}

type UserUseCase interface {
	CreateUser(us User) (*User, error)
}
