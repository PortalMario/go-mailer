package mailer

import (
	"fmt"
	"net/smtp"
)

type Relay struct {
	SMTPHost     string
	SMTPHostPort int
	SMTPUser     string
	SMTPPassword string
}

func NewRelay(SMTPHost, SMTPUser, SMTPPassword string, SMTPHostPort int) *Relay {
	return &Relay{
		SMTPHost:     SMTPHost,
		SMTPHostPort: SMTPHostPort,
		SMTPUser:     SMTPUser,
		SMTPPassword: SMTPPassword,
	}
}

func (relay *Relay) Send(to string, subject string, body string) error {
	msg := "From: " + relay.SMTPUser + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", relay.SMTPUser, relay.SMTPPassword, relay.SMTPHost)

	smtpHostAndPort := fmt.Sprintf("%s:%d", relay.SMTPHost, relay.SMTPHostPort)

	return smtp.SendMail(smtpHostAndPort, auth, relay.SMTPUser, []string{to}, []byte(msg))
}
