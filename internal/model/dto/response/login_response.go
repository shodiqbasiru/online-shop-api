package response

type LoginResponse struct {
	Id    string `json:"id"`
	Role  string `json:"role"`
	Token string `json:"token"`
}
