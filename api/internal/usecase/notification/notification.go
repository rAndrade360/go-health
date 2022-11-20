package notification

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/rAndrade360/go-health/api/internal/domain"
)

type notificationUseCase struct {
	publisher          domain.NotificationPublisher
	notificationSender domain.NotificationSender
}

func NewNotificationUseCase(publisher domain.NotificationPublisher, sender domain.NotificationSender) domain.NotificationUseCase {
	return notificationUseCase{
		publisher:          publisher,
		notificationSender: sender,
	}
}

func (n notificationUseCase) SendToQueue(ntf domain.Notification) error {
	ntf.ID = uuid.NewString()

	b, err := json.Marshal(ntf)
	if err != nil {
		return err
	}

	return n.publisher.Pusblish(b, "application/json")
}

func (n notificationUseCase) Send(nft domain.Notification) error {
	return n.notificationSender.Send(nft)
}
