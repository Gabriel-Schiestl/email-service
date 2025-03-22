package interfaces

import "github.com/Gabriel-Schiestl/email-service/internal/domain/models/message"

type IEmailService[T any] interface {
	SendEmail(content string, msg message.Message[T]) error
}