package request

type CustomerUpdateRequest struct {
	Id           string `validate:"required" json:"id"`
	CustomerName string `validate:"required" json:"customerName"`
	Address      string `validate:"required" json:"address"`
}
