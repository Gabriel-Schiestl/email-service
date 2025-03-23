package usecases

import (
	"encoding/json"

	"github.com/Gabriel-Schiestl/email-service/internal/domain/interfaces"
	"github.com/Gabriel-Schiestl/email-service/internal/domain/models/message"
	"github.com/rabbitmq/amqp091-go"
)

type SendEmailUseCase struct {
	Repository interfaces.ITemplateRepository
	EmailService interfaces.IEmailService
}

func NewSendEmailUseCase(repo interfaces.ITemplateRepository, sendEmailService interfaces.IEmailService) *SendEmailUseCase {
	return &SendEmailUseCase{
		Repository: repo,
		EmailService: sendEmailService,
	}
}

func (s *SendEmailUseCase) Execute(msg amqp091.Delivery) error {
	var message message.Message

	jsonerr := json.Unmarshal(msg.Body, &message)
	if jsonerr != nil {
		return jsonerr
	}

	template, err := s.Repository.GetTemplateById(message.TemplateId)
	if err != nil {
		return err
	}

	senderr := s.EmailService.SendEmail(template.Content, message)
	if senderr != nil {
		return senderr
	}

	return nil
}