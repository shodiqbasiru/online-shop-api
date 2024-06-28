package request

type RegisterRequest struct {
	CustomerName string `validate:"required" json:"customerName"`
	NoHp         string `validate:"required" json:"noHp"`
	Email        string `validate:"required" json:"email"`
	Password     string `validate:"required" json:"password"`
}
