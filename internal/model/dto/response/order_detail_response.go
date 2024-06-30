package response

type OrderDetailResponse struct {
	Id        string `json:"id"`
	Qty       int    `json:"qty"`
	Price     int    `json:"price"`
	OrderId   string `json:"orderId"`
	ProductId string `json:"productId"`
}
