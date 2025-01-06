package client

import "gopkg.in/gomail.v2"

type Smtp struct {
	dialer *gomail.Dialer
}

func NewSmtp(host, from, password string, port int) *Smtp {
	return &Smtp{
		dialer: gomail.NewDialer(host, port, from, password),
	}
}

func (s *Smtp) Send(msg *gomail.Message) error {
	return s.dialer.DialAndSend(msg)
}
