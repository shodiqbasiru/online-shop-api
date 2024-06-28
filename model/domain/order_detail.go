package domain

type OrderDetail struct {
	Id        string
	Qty       int
	Price     int
	OrderId   string
	ProductId string

	Order   Order
	Product Product
}
