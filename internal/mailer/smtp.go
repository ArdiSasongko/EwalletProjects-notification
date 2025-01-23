package mailer

import (
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

const (
	FromName        = "DuckCoding"
	MaxRetires      = 3
	TemplateWelcome = "user_invitations.tmpl"
)

type MailerSMTP interface {
	Send(username, email, subject, body string) error
}

type mailerSMTP struct {
	fromEmail string
	client    *gomail.Dialer
}

func NewClientSMTP(fromEmail, apiKey string) MailerSMTP {
	client := gomail.NewDialer("smtp.gmail.com", 587, fromEmail, apiKey)
	return &mailerSMTP{
		fromEmail: fromEmail,
		client:    client,
	}
}

func (m *mailerSMTP) Send(username, email, subject, body string) error {
	mail := gomail.NewMessage()
	mail.SetAddressHeader("From", m.fromEmail, FromName)
	mail.SetAddressHeader("To", email, username)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	var errRetires error
	for i := 0; i < MaxRetires; i++ {
		if errRetires = m.client.DialAndSend(mail); errRetires != nil {
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		return nil
	}
	return fmt.Errorf("failed to send email after %d attemps, error : %v", MaxRetires, errRetires)
}
