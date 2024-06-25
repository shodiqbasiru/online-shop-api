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

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (controller *ProductController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreateRequest := payload.ProductCreateRequest{}
	err := json.NewDecoder(request.Body).Decode(&productCreateRequest)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.CreateProduct(request.Context(), productCreateRequest)
	commonResponse := response.CommonResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   productResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *ProductController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")

	productResponse := controller.ProductService.GetById(request.Context(), productId)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   productResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *ProductController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productResponses := controller.ProductService.GetAll(request.Context())
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   productResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *ProductController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdateRequest := payload.ProductUpdateRequest{}
	err := json.NewDecoder(request.Body).Decode(&productUpdateRequest)
	helper.PanicIfError(err)

	productId := params.ByName("productId")
	productUpdateRequest.Id = productId

	productResponses := controller.ProductService.UpdateProduct(request.Context(), productUpdateRequest)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   productResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *ProductController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")

	controller.ProductService.DeleteProduct(request.Context(), productId)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}
