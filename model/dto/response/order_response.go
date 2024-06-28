package response

import "time"

type OrderResponse struct {
	Id           string                `json:"id"`
	TransDate    time.Time             `json:"transDate"`
	Status       string                `json:"status"`
	CustomerId   string                `json:"customerId"`
	OrderDetails []OrderDetailResponse `json:"orderDetails"`
}
