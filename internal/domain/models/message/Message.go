package message

type Message struct {
	To         string `json:"to"`
	TemplateId int    `json:"templateId"`
	Subject    string `json:"subject"`
}