package rabbitmq

import (
	"fmt"

	"github.com/rAndrade360/go-health/api/config"
	"github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConection(cfg config.Config) (*amqp091.Connection, error) {
	s := fmt.Sprintf("amqp://%s:%s@%s:%s", cfg.Rabbitmq.User, cfg.Rabbitmq.Password, cfg.Rabbitmq.Host, cfg.Rabbitmq.Port)
	conn, err := amqp091.Dial(s)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
