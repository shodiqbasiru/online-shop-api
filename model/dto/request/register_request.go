package request

type RegisterRequest struct {
	Name     string `json:"name"`
	NoHp     string `json:"noHp"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
