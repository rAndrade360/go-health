package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rAndrade360/go-health/api/adapter/rest/handlers/health"
	"github.com/rAndrade360/go-health/api/adapter/rest/handlers/user"
)

func init() {
	godotenv.Load()
}

func main() {
	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	e.POST("/user", user.Create)
	e.GET("/health", health.Health)
	e.Start(":" + port)
}
