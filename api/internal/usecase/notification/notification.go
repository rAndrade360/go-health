package notification

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/rAndrade360/go-health/api/internal/domain"
)

type notificationUseCase struct {
	publisher domain.NotificationPublisher
}

func NewNotificationUseCase(publisher domain.NotificationPublisher) domain.NotificationUseCase {
	return notificationUseCase{
		publisher: publisher,
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
