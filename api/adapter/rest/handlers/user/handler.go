package user

import (
	echo "github.com/labstack/echo/v4"
	"github.com/rAndrade360/go-health/api/internal/domain"
)

type userHttpHandler struct {
	usecase domain.UserUseCase
}

type UserHttpHandler interface {
	Create(c echo.Context) error
}

func NewUserHttpHandler(usecase domain.UserUseCase) UserHttpHandler {
	return userHttpHandler{
		usecase: usecase,
	}
}

func (u userHttpHandler) Create(c echo.Context) error {
	var userreq domain.User
	err := c.Bind(&userreq)
	if err != nil {
		return err
	}

	us, err := u.usecase.CreateUser(userreq)
	if err != nil {
		return err
	}

	return c.JSON(201, us)
}
