package interfaces

type IEmailService interface {
	SendEmail(content, subject, to string) error
}