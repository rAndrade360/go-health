package user

import (
	"github.com/labstack/echo/v4"
	"github.com/rAndrade360/go-health/api/dto/user"
	"github.com/rAndrade360/go-health/api/internal/model"
	"github.com/rAndrade360/go-health/api/internal/usecase"
)

func Create(c echo.Context) error {
	var userreq user.CreateUserDtoReq
	err := c.Bind(&userreq)
	if err != nil {
		return err
	}

	u := model.User{
		Name:     userreq.Name,
		Username: userreq.Username,
	}

	us, err := usecase.CreateUser(u)
	if err != nil {
		return err
	}

	userres := user.CreateUserDtoRes{
		ID:        us.ID,
		Name:      us.Name,
		Username:  us.Username,
		CreatedAt: us.CreatedAt,
		UpdatedAt: us.UpdatedAt,
	}

	return c.JSON(201, userres)
}
