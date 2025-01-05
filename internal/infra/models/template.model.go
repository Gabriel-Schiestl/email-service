package models

import "gorm.io/gorm"

type Template struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"type:varchar(100)"`
	Content string `gorm:"type:text"`
}

func CriarTabela(db *gorm.DB) error {
	return db.AutoMigrate(&Template{})
}