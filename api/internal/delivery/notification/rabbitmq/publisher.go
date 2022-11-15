package rabbitmq

import (
	"context"

	"github.com/rAndrade360/go-health/api/config"
	"github.com/rAndrade360/go-health/api/internal/domain"
	"github.com/rabbitmq/amqp091-go"
)

var (
	defaultKind = "direct"
)

type rabbitmqPublisher struct {
	amqpChan *amqp091.Channel
	cfg      config.Config
}

func NewNotificationRabbitmqPublisher(cfg config.Config, amqpConn *amqp091.Connection) domain.NotificationPublisher {
	amqpChann, _ := amqpConn.Channel()

	return rabbitmqPublisher{
		amqpChan: amqpChann,
		cfg:      cfg,
	}
}

func (r rabbitmqPublisher) SetupExchangeAndQueue(queueName string) error {
	err := r.amqpChan.ExchangeDeclare(r.cfg.Rabbitmq.Exchange, defaultKind, false, false, false, false, nil)
	if err != nil {
		return err
	}

	q, err := r.amqpChan.QueueDeclare(queueName, false, true, false, false, nil)
	if err != nil {
		return err
	}

	err = r.amqpChan.QueueBind(q.Name, r.cfg.Rabbitmq.RoutingKey, r.cfg.Rabbitmq.Exchange, false, nil)
	if err != nil {
		return err
	}

	return nil
}

func (r rabbitmqPublisher) Pusblish(b []byte, contentType string) error {
	ctx := context.Background()

	return r.amqpChan.PublishWithContext(ctx, r.cfg.Rabbitmq.Exchange, r.cfg.Rabbitmq.RoutingKey, false, false, amqp091.Publishing{
		ContentType: contentType,
		Body:        b,
	})
}
