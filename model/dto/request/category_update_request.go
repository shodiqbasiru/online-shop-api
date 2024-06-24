package request

type CategoryUpdateRequest struct {
	Id   string `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}
