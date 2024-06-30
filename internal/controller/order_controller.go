package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"online-shop-api/internal/helper"
	payload "online-shop-api/internal/model/dto/request"
	"online-shop-api/internal/model/dto/response"
	"online-shop-api/internal/service"
	"online-shop-api/scheduler"
)

type OrderController struct {
	OrderService service.OrderService
	Scheduler    *scheduler.Scheduler
}

func NewOrderController(orderService service.OrderService, scheduler *scheduler.Scheduler) *OrderController {
	return &OrderController{OrderService: orderService, Scheduler: scheduler}
}

func (controller *OrderController) CreateOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderRequest := payload.OrderRequest{}
	err := json.NewDecoder(request.Body).Decode(&orderRequest)
	helper.PanicIfError(err)

	orderResponse := controller.OrderService.CreateOrder(request.Context(), orderRequest)

	controller.Scheduler.ScheduleCancelOrder()

	commonResponse := response.CommonResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   orderResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *OrderController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderId := params.ByName("orderId")

	orderResponse := controller.OrderService.GetById(request.Context(), orderId)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   orderResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *OrderController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderResponses := controller.OrderService.GetAll(request.Context())
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   orderResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *OrderController) UpdateStatusOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderId := params.ByName("orderId")

	orderResponse := controller.OrderService.UpdateStatusOrder(request.Context(), orderId)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   orderResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}
