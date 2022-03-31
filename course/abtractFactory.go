package course

import "fmt"

type InotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Send Notification")
}


type SMSNotificationSender struct{

}


func (SMSNotificationSender) GetSenderMethod () string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel () string {
	return "Twilio"
}



func (SMSNotification) GetSender() ISender {
return SMSNotificationSender{}
}




func AbstractFactory() {

}