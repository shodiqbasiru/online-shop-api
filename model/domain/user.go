package domain

type User struct {
	Id         string
	NoHp       string
	Email      string
	Password   string
	Role       string
	CustomerId string
}

const (
	ADMIN    = "admin"
	CUSTOMER = "customer"
)
