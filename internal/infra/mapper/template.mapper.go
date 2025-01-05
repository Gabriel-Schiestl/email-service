package mapper

import (
	"github.com/Gabriel-Schiestl/email-service/internal/domain/interfaces"
	"github.com/Gabriel-Schiestl/email-service/internal/infra/models"
)

func ModelToDomain(model models.Template) *interfaces.Template {
	return &interfaces.Template{
		ID: model.ID,
		Name: model.Name,
		Content: model.Content,
	}
}