package services

import (
	"bytes"
	"html/template"

	"github.com/Gabriel-Schiestl/email-service/internal/config"
	"github.com/Gabriel-Schiestl/email-service/internal/domain/models/message"
	"gopkg.in/gomail.v2"
)

type EmailService[T any] struct {
	Config *config.SenderConfig
}

func NewEmailService[T any](cfg *config.SenderConfig) *EmailService[T] {
	return &EmailService[T]{
		Config: cfg,
	}
}

func (e *EmailService[T]) SendEmail(content string, msg message.Message[T]) error {
	templ, err := template.New("email").Parse(content)
	if err != nil {
		return err
	}

	var result bytes.Buffer
	if err := templ.Execute(&result, msg.Params); err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.Config.Username)
	mailer.SetHeader("To", msg.To)
	mailer.SetHeader("Subject", msg.Subject)
	mailer.SetBody("text/html", result.String())

	dialer := gomail.NewDialer(e.Config.Host, e.Config.Port, e.Config.Username, e.Config.Password)

	return dialer.DialAndSend(mailer)
}