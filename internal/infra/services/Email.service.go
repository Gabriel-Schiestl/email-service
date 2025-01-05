package services

import (
	"bytes"
	"html/template"

	"github.com/Gabriel-Schiestl/email-service/internal/config"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	Config *config.SenderConfig
}

func NewEmailService(cfg *config.SenderConfig) *EmailService {
	return &EmailService{
		Config: cfg,
	}
}

func (e *EmailService) SendEmail(content, subject, to string) error {

	templ, err := template.New("email").Parse(content)
	if err != nil {
		return err
	}

	var result bytes.Buffer
	if err := templ.Execute(&result, to); err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.Config.Username)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", result.String())

	dialer := gomail.NewDialer(e.Config.Host, e.Config.Port, e.Config.Username, e.Config.Password)

	return dialer.DialAndSend(mailer)
}