package product

type ProductCreateRequest struct {
	Name        string   `json:"name" validate:"required,max=50"`
	Description string   `json:"description" validate:"max=100"`
	Images      []string `json:"images" validate:"dive,url"`
}

type ProductUpdateRequest struct {
	ID          uint
	Name        string   `json:"name" validate:"required,max=50"`
	Description string   `json:"description" validate:"max=100"`
	Images      []string `json:"images" validate:"dive,url"`
}
