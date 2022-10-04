package user

import "time"

type CreateUserDtoReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type CreateUserDtoRes struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
