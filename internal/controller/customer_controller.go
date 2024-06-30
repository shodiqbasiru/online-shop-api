package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"online-shop-api/internal/helper"
	payload "online-shop-api/internal/model/dto/request"
	"online-shop-api/internal/model/dto/response"
	"online-shop-api/internal/service"
)

type CustomerController struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) *CustomerController {
	return &CustomerController{CustomerService: customerService}
}

func (controller *CustomerController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")

	customerResponse := controller.CustomerService.GetById(request.Context(), customerId)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   customerResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *CustomerController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerResponse := controller.CustomerService.GetAll(request.Context())
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   customerResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *CustomerController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerUpdateRequest := payload.CustomerUpdateRequest{}
	err := json.NewDecoder(request.Body).Decode(&customerUpdateRequest)
	helper.PanicIfError(err)

	customerId := params.ByName("customerId")
	customerUpdateRequest.Id = customerId

	customerResponse := controller.CustomerService.UpdateCustomer(request.Context(), customerUpdateRequest)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   customerResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *CustomerController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")

	controller.CustomerService.DeleteCustomer(request.Context(), customerId)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}
