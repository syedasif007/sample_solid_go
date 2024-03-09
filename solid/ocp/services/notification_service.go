package services

type NotificationService interface {
	SendNotification(message string) error
}
