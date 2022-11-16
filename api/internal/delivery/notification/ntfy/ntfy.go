package ntfy

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rAndrade360/go-health/api/internal/domain"
)

func Send(ntf domain.Notification) error {
	url := fmt.Sprintf("https://ntfy.sh/%s", ntf.UserID)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, strings.NewReader(ntf.Message))
	if err != nil {
		return err
	}

	req.Header.Set("Title", ntf.Title)

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
