package message

type Message[T any] struct {
	To         string `json:"to"`
	TemplateId int    `json:"templateId"`
	Subject    string `json:"subject"`
	Params     T      `json:"params"`
}