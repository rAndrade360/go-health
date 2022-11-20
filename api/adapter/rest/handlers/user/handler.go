package user

import (
	echo "github.com/labstack/echo/v4"
	"github.com/rAndrade360/go-health/api/internal/domain"
)

type userHttpHandler struct {
	usecase             domain.UserUseCase
	notificationusecase domain.NotificationUseCase
}

type UserHttpHandler interface {
	Create(c echo.Context) error
}

func NewUserHttpHandler(usecase domain.UserUseCase, notificationusecase domain.NotificationUseCase) UserHttpHandler {
	return userHttpHandler{
		usecase:             usecase,
		notificationusecase: notificationusecase,
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

	err = u.notificationusecase.SendToQueue(domain.Notification{
		Title:   "User Creation",
		Message: "User Created Successfully",
		UserID:  us.Username,
		Via:     "ntfy",
	})
	if err != nil {
		return err
	}

	return c.JSON(201, us)
}
