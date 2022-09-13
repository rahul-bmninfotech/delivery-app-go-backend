package utils

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

// SendMail mails the file to the emailID
func SendMail(file, filename, emailID string) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "deliveryapp001@gmail.com")
	mail.SetHeader("To", emailID)
	mail.SetHeader("Subject", fmt.Sprintf("Comment from Buyer %s", filename))
	mail.SetBody("text/plain", "")

	mail.Attach(file)
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "deliveryapp001@gmail.com", "DeliveryApp@1")
	if err := dialer.DialAndSend(mail); err != nil {
		log.Println(err)
	}
}
