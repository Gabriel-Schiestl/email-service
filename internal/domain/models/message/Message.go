package message

type Message struct {
	Pattern string `json:"pattern"`
	Data    struct {
		To         string                 `json:"to"`
		TemplateId int                    `json:"templateId"`
		Subject    string                 `json:"subject"`
		Params     map[string]interface{} `json:"params"`
	} `json:"data"`
}