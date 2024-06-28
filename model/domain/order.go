package domain

import "time"

type Order struct {
	Id         string
	Status     string
	TransDate  time.Time
	CustomerId string

	Customer     Customer
	OrderDetails []OrderDetail
}
