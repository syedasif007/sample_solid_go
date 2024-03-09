package impls

import (
	"fmt"
)

type SMSNotificationService struct{}

func (sns *SMSNotificationService) SendNotification(message string) error {
	fmt.Println("Notification:", message)
	return nil
}
