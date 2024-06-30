package request

type OrderRequest struct {
	CustomerId   string               `json:"customerId"`
	OrderDetails []OrderDetailRequest `json:"orderDetails"`
}
