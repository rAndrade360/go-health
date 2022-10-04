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

	err = usecase.CreateUser(&u)
	if err != nil {
		return err
	}

	userres := user.CreateUserDtoRes{
		ID:        u.ID,
		Name:      u.Name,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	return c.JSON(200, userres)
}
