package request

type OrderDetailRequest struct {
	ProductId string `json:"productId"`
	Qty       int    `json:"qty"`
}
