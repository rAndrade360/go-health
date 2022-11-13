package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rAndrade360/go-health/api/adapter/rest/handlers/health"
	userhttp "github.com/rAndrade360/go-health/api/adapter/rest/handlers/user"
	userrepo "github.com/rAndrade360/go-health/api/internal/repository/user"
	userusecase "github.com/rAndrade360/go-health/api/internal/usecase"
)

func init() {
	godotenv.Load()
}

func main() {
	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := sql.DB{}

	userRepo := userrepo.NewUserRepository(&db)
	userUseCase := userusecase.NewUserUseCase(userRepo)
	userHttpHandler := userhttp.NewUserHttpHandler(userUseCase)

	e.POST("/user", userHttpHandler.Create)
	e.GET("/health", health.Health)
	log.Fatal(e.Start(":" + port))
}
