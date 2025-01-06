package service

import (
	"gopkg.in/gomail.v2"
	"notifier-service/client"
	"notifier-service/domain"
)

type Sender interface {
	Send(message domain.Mail) error
}

type SenderImpl struct {
	receiver   string
	sender     string
	smtpClient *client.Smtp
}

func NewSender(receiver string, sender string, smtpClient *client.Smtp) Sender {
	return &SenderImpl{
		receiver: receiver,
	}
}

func (s *SenderImpl) Send(message domain.Mail) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", s.sender)
	msg.SetHeader("To", s.receiver)
	msg.SetHeader("Subject", message.Subject)
	msg.SetBody("text/plain", message.Body)

	return s.smtpClient.Send(msg)
}
