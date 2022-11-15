package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rAndrade360/go-health/api/adapter/database/sqlite"
	"github.com/rAndrade360/go-health/api/adapter/rabbitmq"
	"github.com/rAndrade360/go-health/api/adapter/rest/handlers/health"
	userhttp "github.com/rAndrade360/go-health/api/adapter/rest/handlers/user"
	"github.com/rAndrade360/go-health/api/config"
	notificationrabbit "github.com/rAndrade360/go-health/api/internal/delivery/notification/rabbitmq"
	userrepo "github.com/rAndrade360/go-health/api/internal/repository/user/sqlite"
	"github.com/rAndrade360/go-health/api/internal/usecase/notification"
	userusecase "github.com/rAndrade360/go-health/api/internal/usecase/user"
)

func init() {
	godotenv.Load("../../.env")
}

func main() {
	cgf := config.LoadConfig()
	conn, err := rabbitmq.NewRabbitMQConection(cgf)
	if err != nil {
		log.Fatal("Err to start rabbit: ", err.Error())
	}

	npublisher := notificationrabbit.NewNotificationRabbitmqPublisher(cgf, conn)
	npublisher.SetupExchangeAndQueue("nty_queue")

	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := sqlite.GetDB()
	if err != nil {
		log.Fatalf("Err to connect to DB: %s", err.Error())
	}

	userRepo := userrepo.NewUserSqliteRepository(db)

	userUseCase := userusecase.NewUserUseCase(userRepo)
	notificationUseCase := notification.NewNotificationUseCase(npublisher)

	userHttpHandler := userhttp.NewUserHttpHandler(userUseCase, notificationUseCase)

	e.POST("/user", userHttpHandler.Create)
	e.GET("/health", health.Health)
	log.Fatal(e.Start(":" + port))
}
