package usecases

import (
	"encoding/json"

	"github.com/Gabriel-Schiestl/email-service/internal/domain/interfaces"
	"github.com/Gabriel-Schiestl/email-service/internal/domain/models/message"
	"github.com/rabbitmq/amqp091-go"
)

type SendEmailUseCase[T any] struct {
	Repository interfaces.ITemplateRepository
	EmailService interfaces.IEmailService[T]
}

func NewSendEmailUseCase[T any](repo interfaces.ITemplateRepository, sendEmailService interfaces.IEmailService[T]) *SendEmailUseCase[T] {
	return &SendEmailUseCase[T]{
		Repository: repo,
		EmailService: sendEmailService,
	}
}

func (s *SendEmailUseCase[T]) Execute(msg amqp091.Delivery) error {
	var message message.Message[T]

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