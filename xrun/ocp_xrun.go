package xrun

import (
	"fmt"
	"sample_solid_go/solid/ocp/impls"
	"sample_solid_go/solid/ocp/utils"
)

func OCP() {

	fmt.Println("Result of OCP:")

	emailNfService := &impls.EmailNotificationService{}
	nfSender := &utils.NotificationSender{NotificationService: emailNfService}
	nfSender.SendNotification("Email")

	nfSender.NotificationService = &impls.SMSNotificationService{}
	nfSender.SendNotification("SMS")

	fmt.Println()
}
