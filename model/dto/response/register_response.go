package response

type RegisterResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	NoHp  string `json:"noHp"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
