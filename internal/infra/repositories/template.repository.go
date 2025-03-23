package repositories

import (
	"errors"

	"github.com/Gabriel-Schiestl/email-service/internal/domain/interfaces"
	"github.com/Gabriel-Schiestl/email-service/internal/infra/mapper"
	"github.com/Gabriel-Schiestl/email-service/internal/infra/models"
	"gorm.io/gorm"
)

type TemplateRepository struct {
	DB *gorm.DB
}

func NewTemplateRepository(db *gorm.DB) *TemplateRepository {
	return &TemplateRepository{
		DB: db,
	}
}

func (t *TemplateRepository) GetTemplateById(id int) (*interfaces.Template, error) {

	var template models.Template
	if err := t.DB.First(&template, id).Error; err != nil {
		return nil, errors.New("template not found")
	}

	return mapper.ModelToDomain(template), nil
}
