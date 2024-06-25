package request

type ProductCreateRequest struct {
	Name        string `validate:"required,max=255,min=1" json:"name"`
	Description string `validate:"required,min=1" json:"description"`
	Price       int    `validate:"required min:1" json:"price"`
	Stock       int    `validate:"required min=1" json:"stock"`
	CategoryId  string `validate:"required" json:"categoryId"`
}
