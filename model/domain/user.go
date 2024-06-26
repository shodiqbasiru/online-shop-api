package domain

type User struct {
	Id         string
	NoHp       string
	Email      string
	Password   string
	Role       Role
	CustomerId string
}

type Role int

const (
	ADMIN Role = iota
	USER
)

func (r Role) String() string {
	return [...]string{"ADMIN", "USER"}[r]
}
