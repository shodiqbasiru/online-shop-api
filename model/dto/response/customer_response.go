package response

type CustomerResponse struct {
	Id           string `json:"id"`
	CustomerName string `json:"customerName"`
	Address      string `json:"address"`
}
