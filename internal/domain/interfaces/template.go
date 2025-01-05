package interfaces

type Template struct {
	ID      int
	Name    string
	Content string
}

type ITemplateRepository interface {
	GetTemplateById(id int) (*Template, error)
}