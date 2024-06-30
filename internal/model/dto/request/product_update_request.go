package request

type ProductUpdateRequest struct {
	Id          string `validate:"required"`
	Name        string `validate:"required,max=255,min=1" json:"name"`
	Description string `validate:"required,min=1" json:"description"`
	Price       int    `validate:"gt=1" json:"price"`
	Stock       int    `validate:"gt=1" json:"stock"`
	CategoryId  string `validate:"required" json:"categoryId"`
}
