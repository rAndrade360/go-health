package domain

type Notification struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
	UserID  string `json:"userId"`
	Via     string `json:"via"`
}

type NotificationUseCase interface {
	SendToQueue(ntf Notification) error
}

type NotificationPublisher interface {
	SetupExchangeAndQueue(queueName string) error
	Pusblish(b []byte, contentType string) error
}
