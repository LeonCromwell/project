package email

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(subject string, body string, to []string, cc []string, bcc []string) error
}

type GmailSender struct {
	Username          string
	fromEmailAdress   string
	fromEmailPassword string
}

func NewGmailSender(Username, fromEmailUsername, fromEmailPassword string) *GmailSender {
	return &GmailSender{
		Username:          Username,
		fromEmailAdress:   fromEmailUsername,
		fromEmailPassword: fromEmailPassword,
	}
}

func (g *GmailSender) SendEmail(subject string, body string, to []string, cc []string, bcc []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", g.Username, g.fromEmailAdress)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	e.HTML = []byte(body)
	e.Text = []byte(body)
	plainAuth := smtp.PlainAuth("", g.fromEmailAdress, g.fromEmailPassword, smtpAuthAddress)
	err := e.Send(smtpServerAddress, plainAuth)
	if err != nil {
		return err
	}
	return nil
}
