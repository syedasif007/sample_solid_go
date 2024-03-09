package utils

import "sample_solid_go/solid/ocp/services"

type NotificationSender struct {
	NotificationService services.NotificationService
}

func (ns *NotificationSender) SendNotification(message string) error {
	return ns.NotificationService.SendNotification(message)
}
