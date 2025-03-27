package services

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/Gabriel-Schiestl/email-service/internal/config"
	"github.com/Gabriel-Schiestl/email-service/internal/domain/models/message"
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

func (e *EmailService) SendEmail(content string, msg message.Message) error {
	templ, err := template.New("email").Funcs(template.FuncMap{
		"param": func(params map[string]interface{}, key string) string {
			if val, ok := params[key]; ok {
				return fmt.Sprintf("%v", val)
			}
			return ""
		},
	}).Parse(content)
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