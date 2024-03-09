package impls

import (
	"fmt"
)

type EmailNotificationService struct{}

func (ens *EmailNotificationService) SendNotification(message string) error {
	fmt.Println("Notification:", message)
	return nil
}
