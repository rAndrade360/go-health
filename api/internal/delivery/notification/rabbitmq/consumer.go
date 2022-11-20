package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/rAndrade360/go-health/api/config"
	"github.com/rAndrade360/go-health/api/internal/domain"
	"github.com/rabbitmq/amqp091-go"
)

type rabbitmqConsumer struct {
	amqpChan            *amqp091.Channel
	queue               *amqp091.Queue
	cfg                 config.Config
	notificationUseCase domain.NotificationUseCase
}

func NewNotificationRabbitmqConsumer(cfg config.Config, amqpConn *amqp091.Connection, notificationUseCase domain.NotificationUseCase) domain.NotificationConsumer {
	amqpChann, err := amqpConn.Channel()
	if err != nil {
		log.Println("Error to create consumer chann: ", err.Error())
	}

	return rabbitmqConsumer{
		amqpChan:            amqpChann,
		cfg:                 cfg,
		notificationUseCase: notificationUseCase,
	}
}

func (r rabbitmqConsumer) SetupExchangeAndQueue(queueName string) error {
	err := r.amqpChan.ExchangeDeclare(r.cfg.Rabbitmq.Exchange, defaultKind, false, false, false, false, nil)
	if err != nil {
		return err
	}

	q, err := r.amqpChan.QueueDeclare(queueName, false, true, false, false, nil)
	if err != nil {
		return err
	}

	err = r.amqpChan.QueueBind(q.Name, r.cfg.Rabbitmq.RoutingKey, r.cfg.Rabbitmq.Exchange, true, nil)
	if err != nil {
		return err
	}

	return nil
}

func (r rabbitmqConsumer) Consume(queueName string) error {
	msgs, err := r.amqpChan.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var n domain.Notification

			err = json.Unmarshal(d.Body, &n)
			if err != nil {
				log.Println("err to consume ", err.Error())
				d.Ack(false)
				continue
			}

			r.notificationUseCase.Send(n)
			d.Ack(false)
		}
	}()

	<-forever

	return nil
}
