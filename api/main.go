package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rAndrade360/go-health/api/adapter/rest/handlers/user"
)

func main() {
	e := echo.New()

	e.POST("/user", user.Create)
}
