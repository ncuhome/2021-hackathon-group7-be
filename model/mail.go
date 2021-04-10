package model

import (
	"github.com/go-gomail/gomail"
)

type GoMail struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func (s *GoMail) Init(host string, port int, username string, password string, name string) {
	s.Host = host
	s.Port = port
	s.Username = username
	s.Password = password
	s.Name = name
	return
}

func (s *GoMail) Send(toEmail string, toName string, subject string, contentType string, content string) error {
	message := gomail.NewMessage()
	message.SetAddressHeader("From", s.Username, s.Name)
	message.SetHeader("To", message.FormatAddress(toEmail, toName))
	message.SetHeader("Subject", subject)
	message.SetBody(contentType, content)

	err := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password).DialAndSend(message)
	if err != nil {
		return err
	}

	return nil
}
