package request

type LoginRequest struct {
	EmailOrNoHp string `json:"emailOrNoHp"`
	Password    string `json:"password"`
}
