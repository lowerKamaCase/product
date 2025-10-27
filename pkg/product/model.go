package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `json:"name" validate:"max=50"`
	Description string   `json:"description" validate:"max=100"`
	Images      []string `json:"images" gorm:"type:text[]"`
}

func NewProduct(name string, description string, images []string) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Images:      images,
	}
}
