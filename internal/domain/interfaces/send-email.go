package interfaces

import "github.com/Gabriel-Schiestl/email-service/internal/domain/models/message"

type IEmailService interface {
	SendEmail(content string, msg message.Message) error
}