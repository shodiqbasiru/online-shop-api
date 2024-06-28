package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"online-shop-api/helper"
	payload "online-shop-api/model/dto/request"
	"online-shop-api/model/dto/response"
	"online-shop-api/service"
)

type OrderController struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}

func (controller *OrderController) CreateOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderRequest := payload.OrderRequest{}
	err := json.NewDecoder(request.Body).Decode(&orderRequest)
	helper.PanicIfError(err)

	orderResponse := controller.OrderService.CreateOrder(request.Context(), orderRequest)
	commonResponse := response.CommonResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   orderResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}
